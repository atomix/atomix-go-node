// Copyright 2019-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package time

import (
	"sync"
	"time"
)

// Clock is an interface for clocks
type Clock interface {
	// Get gets the current timestamp
	Get() Timestamp
	// Increment increments the clock
	Increment() Timestamp
	// Update updates the timestamp
	Update(Timestamp) Timestamp
}

// NewPhysicalClock creates a new physical clock
func NewPhysicalClock() Clock {
	return &PhysicalClock{
		timestamp: NewPhysicalTimestamp(PhysicalTime(time.Now())),
	}
}

// PhysicalClock is a clock that produces PhysicalTimestamps
type PhysicalClock struct {
	timestamp Timestamp
	mu        sync.RWMutex
}

func (c *PhysicalClock) Get() Timestamp {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.timestamp
}

func (c *PhysicalClock) Increment() Timestamp {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.timestamp = NewPhysicalTimestamp(PhysicalTime(time.Now()))
	return c.timestamp
}

func (c *PhysicalClock) Update(update Timestamp) Timestamp {
	c.mu.RLock()
	current := c.timestamp
	c.mu.RUnlock()
	if !update.After(current) {
		return current
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	if update.After(c.timestamp) {
		c.timestamp = update
	}
	return c.timestamp
}

// NewLogicalClock creates a new logical clock
func NewLogicalClock() Clock {
	return &LogicalClock{
		timestamp: NewLogicalTimestamp(LogicalTime(0)).(LogicalTimestamp),
	}
}

// LogicalClock is a clock that produces LogicalTimestamps
type LogicalClock struct {
	timestamp LogicalTimestamp
	mu        sync.RWMutex
}

func (c *LogicalClock) Get() Timestamp {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.timestamp
}

func (c *LogicalClock) Increment() Timestamp {
	c.mu.Lock()
	defer c.mu.Unlock()
	timestamp := LogicalTimestamp{
		Time: c.timestamp.Time + 1,
	}
	c.timestamp = timestamp
	return timestamp
}

func (c *LogicalClock) Update(t Timestamp) Timestamp {
	update, ok := t.(LogicalTimestamp)
	if !ok {
		panic("not a logical timestamp")
	}

	c.mu.RLock()
	current := c.timestamp
	c.mu.RUnlock()
	if !update.After(current) {
		return current
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	if update.After(c.timestamp) {
		c.timestamp = update
	}
	return c.timestamp
}

// NewCompositeClock creates a new composite clock
func NewCompositeClock(clocks ...Clock) Clock {
	timestamps := make([]Timestamp, len(clocks))
	for i, clock := range clocks {
		timestamps[i] = clock.Get()
	}
	return &CompositeClock{
		clocks:    clocks,
		timestamp: NewCompositeTimestamp(timestamps...),
	}
}

// CompositeClock is a clock that produces CompositeTimestamps
type CompositeClock struct {
	clocks    []Clock
	timestamp Timestamp
	mu        sync.RWMutex
}

func (c *CompositeClock) Get() Timestamp {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.timestamp
}

func (c *CompositeClock) Increment() Timestamp {
	timestamps := make([]Timestamp, len(c.clocks))
	for i, clock := range c.clocks {
		timestamps[i] = clock.Increment()
	}
	return NewCompositeTimestamp(timestamps...)
}

func (c *CompositeClock) Update(t Timestamp) Timestamp {
	update, ok := t.(CompositeTimestamp)
	if !ok {
		panic("not a composite timestamp")
	}

	timestamps := make([]Timestamp, len(c.clocks))
	for i, clock := range c.clocks {
		timestamps[i] = clock.Update(update.Timestamps[i])
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.timestamp = NewCompositeTimestamp(timestamps...)
	return c.timestamp
}
