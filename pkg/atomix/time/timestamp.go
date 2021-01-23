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
	metaapi "github.com/atomix/api/go/atomix/primitive/meta"
	"time"
)

// NewTimestamp creates new object timestamp from the given proto timestamp
func NewTimestamp(meta metaapi.Timestamp) Timestamp {
	switch t := meta.Timestamp.(type) {
	case *metaapi.Timestamp_PhysicalTimestamp:
		return NewPhysicalTimestamp(PhysicalTime(t.PhysicalTimestamp.Time))
	case *metaapi.Timestamp_LogicalTimestamp:
		return NewLogicalTimestamp(LogicalTime(t.LogicalTimestamp.Time))
	case *metaapi.Timestamp_VectorTimestamp:
		times := make([]LogicalTime, len(t.VectorTimestamp.Time))
		for i, time := range t.VectorTimestamp.Time {
			times[i] = LogicalTime(time)
		}
		return NewVectorTimestamp(times, 0)
	case *metaapi.Timestamp_EpochTimestamp:
		return NewEpochTimestamp(Epoch(t.EpochTimestamp.Epoch.Num), LogicalTime(t.EpochTimestamp.Sequence.Num))
	default:
		panic("unknown timestamp type")
	}
}

// Timestamp is a request timestamp
type Timestamp interface {
	Before(Timestamp) bool
	After(Timestamp) bool
	Equal(Timestamp) bool
	Proto() metaapi.Timestamp
}

type LogicalTime uint64

func NewLogicalTimestamp(time LogicalTime) Timestamp {
	return LogicalTimestamp{
		Time: time,
	}
}

type LogicalTimestamp struct {
	Time LogicalTime
}

func (t LogicalTimestamp) Increment() LogicalTimestamp {
	return LogicalTimestamp{
		Time: t.Time + 1,
	}
}

func (t LogicalTimestamp) Before(u Timestamp) bool {
	v, ok := u.(LogicalTimestamp)
	if !ok {
		panic("not a logical timestamp")
	}
	return t.Time < v.Time
}

func (t LogicalTimestamp) After(u Timestamp) bool {
	v, ok := u.(LogicalTimestamp)
	if !ok {
		panic("not a logical timestamp")
	}
	return t.Time > v.Time
}

func (t LogicalTimestamp) Equal(u Timestamp) bool {
	v, ok := u.(LogicalTimestamp)
	if !ok {
		panic("not a logical timestamp")
	}
	return t.Time == v.Time
}

func (t LogicalTimestamp) Proto() metaapi.Timestamp {
	return metaapi.Timestamp{
		Timestamp: &metaapi.Timestamp_LogicalTimestamp{
			LogicalTimestamp: &metaapi.LogicalTimestamp{
				Time: metaapi.LogicalTime(t.Time),
			},
		},
	}
}

func NewVectorTimestamp(times []LogicalTime, i int) Timestamp {
	return VectorTimestamp{
		Times: times,
		i:     i,
	}
}

type VectorTimestamp struct {
	Times []LogicalTime
	i     int
}

func (t VectorTimestamp) Before(u Timestamp) bool {
	v, ok := u.(VectorTimestamp)
	if !ok {
		panic("not a vector timestamp")
	}
	for i := range v.Times {
		if t.Times[t.i] >= v.Times[i] {
			return false
		}
	}
	return true
}

func (t VectorTimestamp) After(u Timestamp) bool {
	v, ok := u.(VectorTimestamp)
	if !ok {
		panic("not a vector timestamp")
	}
	for i := range v.Times {
		if t.Times[t.i] <= v.Times[i] {
			return false
		}
	}
	return true
}

func (t VectorTimestamp) Equal(u Timestamp) bool {
	v, ok := u.(VectorTimestamp)
	if !ok {
		panic("not a vector timestamp")
	}
	for i := range v.Times {
		if t.Times[t.i] != v.Times[i] {
			return false
		}
	}
	return true
}

func (t VectorTimestamp) Proto() metaapi.Timestamp {
	times := make([]metaapi.LogicalTime, len(t.Times))
	for i, time := range t.Times {
		times[i] = metaapi.LogicalTime(time)
	}
	return metaapi.Timestamp{
		Timestamp: &metaapi.Timestamp_VectorTimestamp{
			VectorTimestamp: &metaapi.VectorTimestamp{
				Time: times,
			},
		},
	}
}

type PhysicalTime time.Time

func NewPhysicalTimestamp(time PhysicalTime) Timestamp {
	return PhysicalTimestamp{
		Time: time,
	}
}

type PhysicalTimestamp struct {
	Time PhysicalTime
}

func (t PhysicalTimestamp) Before(u Timestamp) bool {
	v, ok := u.(PhysicalTimestamp)
	if !ok {
		panic("not a wall clock timestamp")
	}
	return time.Time(t.Time).Before(time.Time(v.Time))
}

func (t PhysicalTimestamp) After(u Timestamp) bool {
	v, ok := u.(PhysicalTimestamp)
	if !ok {
		panic("not a wall clock timestamp")
	}
	return time.Time(t.Time).After(time.Time(v.Time))
}

func (t PhysicalTimestamp) Equal(u Timestamp) bool {
	v, ok := u.(PhysicalTimestamp)
	if !ok {
		panic("not a wall clock timestamp")
	}
	return time.Time(t.Time).Equal(time.Time(v.Time))
}

func (t PhysicalTimestamp) Proto() metaapi.Timestamp {
	return metaapi.Timestamp{
		Timestamp: &metaapi.Timestamp_PhysicalTimestamp{
			PhysicalTimestamp: &metaapi.PhysicalTimestamp{
				Time: metaapi.PhysicalTime(t.Time),
			},
		},
	}
}

type Epoch uint64

func NewEpochTimestamp(epoch Epoch, time LogicalTime) Timestamp {
	return EpochTimestamp{
		Epoch: epoch,
		Time:  time,
	}
}

type EpochTimestamp struct {
	Epoch Epoch
	Time  LogicalTime
}

func (t EpochTimestamp) Before(u Timestamp) bool {
	v, ok := u.(EpochTimestamp)
	if !ok {
		panic("not an epoch timestamp")
	}
	return t.Epoch < v.Epoch || (t.Epoch == v.Epoch && t.Time < v.Time)
}

func (t EpochTimestamp) After(u Timestamp) bool {
	v, ok := u.(EpochTimestamp)
	if !ok {
		panic("not an epoch timestamp")
	}
	return t.Epoch > v.Epoch || (t.Epoch == v.Epoch && t.Time > v.Time)
}

func (t EpochTimestamp) Equal(u Timestamp) bool {
	v, ok := u.(EpochTimestamp)
	if !ok {
		panic("not an epoch timestamp")
	}
	return t.Epoch == v.Epoch && t.Time == v.Time
}

func (t EpochTimestamp) Proto() metaapi.Timestamp {
	return metaapi.Timestamp{
		Timestamp: &metaapi.Timestamp_EpochTimestamp{
			EpochTimestamp: &metaapi.EpochTimestamp{
				Epoch: metaapi.Epoch{
					Num: metaapi.EpochNum(t.Epoch),
				},
				Sequence: metaapi.Sequence{
					Num: metaapi.SequenceNum(t.Time),
				},
			},
		},
	}
}

func NewCompositeTimestamp(timestamps ...Timestamp) Timestamp {
	return CompositeTimestamp{
		Timestamps: timestamps,
	}
}

type CompositeTimestamp struct {
	Timestamps []Timestamp
}

func (t CompositeTimestamp) Before(u Timestamp) bool {
	v, ok := u.(CompositeTimestamp)
	if !ok {
		panic("not a composite timestamp")
	}
	if len(t.Timestamps) != len(v.Timestamps) {
		panic("incompatible composite timestamps")
	}
	for i := 0; i < len(t.Timestamps); i++ {
		t1 := t.Timestamps[i]
		t2 := v.Timestamps[i]
		if t1.Before(t2) {
			return true
		} else if i > 0 {
			for j := 0; j < i; j++ {
				v1 := t.Timestamps[j]
				v2 := v.Timestamps[j]
				if !v1.Equal(v2) {
					return false
				}
			}
			if !t1.Before(t2) {
				return false
			}
		}
	}
	return true
}

func (t CompositeTimestamp) After(u Timestamp) bool {
	v, ok := u.(CompositeTimestamp)
	if !ok {
		panic("not a composite timestamp")
	}
	if len(t.Timestamps) != len(v.Timestamps) {
		panic("incompatible composite timestamps")
	}
	for i := 0; i < len(t.Timestamps); i++ {
		t1 := t.Timestamps[i]
		t2 := v.Timestamps[i]
		if t1.After(t2) {
			return true
		} else if i > 0 {
			for j := 0; j < i; j++ {
				v1 := t.Timestamps[j]
				v2 := v.Timestamps[j]
				if !v1.Equal(v2) {
					return false
				}
			}
			if !t1.After(t2) {
				return false
			}
		}
	}
	return true
}

func (t CompositeTimestamp) Equal(u Timestamp) bool {
	v, ok := u.(CompositeTimestamp)
	if !ok {
		panic("not a composite timestamp")
	}
	if len(t.Timestamps) != len(v.Timestamps) {
		panic("incompatible composite timestamps")
	}
	for i := 0; i < len(t.Timestamps); i++ {
		t1 := t.Timestamps[i]
		t2 := v.Timestamps[i]
		if !t1.Equal(t2) {
			return false
		}
	}
	return true
}

func (t CompositeTimestamp) Proto() metaapi.Timestamp {
	timestamps := make([]metaapi.Timestamp, 0, len(t.Timestamps))
	for _, timestamp := range t.Timestamps {
		timestamps = append(timestamps, timestamp.Proto())
	}
	return metaapi.Timestamp{
		Timestamp: &metaapi.Timestamp_CompositeTimestamp{
			CompositeTimestamp: &metaapi.CompositeTimestamp{
				Timestamps: timestamps,
			},
		},
	}
}
