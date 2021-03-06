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

package cluster

import (
	"context"
	protocolapi "github.com/atomix/atomix-api/go/atomix/protocol"
	"github.com/atomix/atomix-go-framework/pkg/atomix/errors"
	"sync"
)

// PartitionID is a partition identifier
type PartitionID uint32

// NewPartition returns a new replica
func NewPartition(config protocolapi.ProtocolPartition, cluster Cluster) Partition {
	return &partition{
		id:           PartitionID(config.PartitionID),
		cluster:      cluster,
		replicasByID: make(map[ReplicaID]*Replica),
	}
}

// Partition is a cluster partition
type Partition interface {
	// ID returns the partition identifier
	ID() PartitionID
	// Member returns the local partition member
	Member() (*Member, bool)
	// Replica looks up a replica in the partition
	Replica(id ReplicaID) (*Replica, bool)
	// Replicas returns the set of all replicas in the partition
	Replicas() []*Replica
	// Watch watches the partition for changes
	Watch(ctx context.Context, ch chan<- ReplicaSet) error
}

// ConfigurablePartition is an interface for configurable Partitions
type ConfigurablePartition interface {
	Update(protocolapi.ProtocolPartition) error
}

// partition is a cluster partition
type partition struct {
	id           PartitionID
	cluster      Cluster
	replicas     ReplicaSet
	replicasByID map[ReplicaID]*Replica
	watchers     []chan<- ReplicaSet
	configMu     sync.RWMutex
	updateMu     sync.Mutex
}

func (p *partition) ID() PartitionID {
	return p.id
}

// Member returns the local partition member
func (p *partition) Member() (*Member, bool) {
	member, ok := p.cluster.Member()
	if !ok {
		return nil, false
	}
	_, ok = p.Replica(member.ID)
	if !ok {
		return nil, false
	}
	return member, true
}

// Replica returns a replica by ID
func (p *partition) Replica(id ReplicaID) (*Replica, bool) {
	replica, ok := p.replicasByID[id]
	return replica, ok
}

// Replicas returns the current replicas
func (p *partition) Replicas() []*Replica {
	p.configMu.RLock()
	defer p.configMu.RUnlock()
	replicas := make([]*Replica, len(p.replicas))
	copy(replicas, p.replicas)
	return replicas
}

// Update updates the partition configuration
func (p *partition) Update(config protocolapi.ProtocolPartition) error {
	p.updateMu.Lock()
	defer p.updateMu.Unlock()
	p.configMu.Lock()

	replicas := make(ReplicaSet, 0, len(config.Replicas))
	replicasByID := make(map[ReplicaID]*Replica)
	for _, id := range config.Replicas {
		replicaID := ReplicaID(id)
		replica, ok := p.cluster.Replica(replicaID)
		if !ok {
			return errors.NewNotFound("replica '%s' not a member of the cluster", replicaID)
		}
		replicas = append(replicas, replica)
		replicasByID[replicaID] = replica
	}

	for id := range p.replicasByID {
		if _, ok := replicasByID[id]; !ok {
			delete(p.replicasByID, id)
		}
	}

	for id, replica := range replicasByID {
		if _, ok := p.replicasByID[id]; !ok {
			p.replicasByID[id] = replica
		}
	}
	p.replicas = replicas
	p.configMu.Unlock()

	p.configMu.RLock()
	for _, watcher := range p.watchers {
		watcher <- replicas
	}
	p.configMu.RUnlock()
	return nil
}

// Watch watches the partition for changes
func (p *partition) Watch(ctx context.Context, ch chan<- ReplicaSet) error {
	p.configMu.Lock()
	watcher := make(chan ReplicaSet)
	replicas := p.replicas
	go func() {
		if replicas != nil {
			ch <- replicas
		}
		for {
			select {
			case replicas, ok := <-watcher:
				if !ok {
					return
				}
				ch <- replicas
			case <-ctx.Done():
				p.configMu.Lock()
				watchers := make([]chan<- ReplicaSet, 0)
				for _, ch := range p.watchers {
					if ch != watcher {
						watchers = append(watchers, ch)
					}
				}
				p.watchers = watchers
				p.configMu.Unlock()
				close(watcher)
			}
		}
	}()
	p.watchers = append(p.watchers, watcher)
	p.configMu.Unlock()
	return nil
}

var _ Partition = &partition{}
var _ ConfigurablePartition = &partition{}
