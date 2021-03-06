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

package gossip

import (
	"context"
	"github.com/atomix/atomix-api/go/atomix/primitive/meta"
	"github.com/atomix/atomix-go-framework/pkg/atomix/cluster"
	"github.com/atomix/atomix-go-framework/pkg/atomix/errors"
	"sync"
)

// NewPartition creates a new proxy partition
func NewPartition(p cluster.Partition, registry *Registry) *Partition {
	return &Partition{
		ID:        PartitionID(p.ID()),
		Partition: p,
		registry:  registry,
		services:  make(map[ServiceId]Service),
		replicas:  make(map[ServiceId]Replica),
	}
}

// PartitionID is a partition identifier
type PartitionID int

// Partition is a proxy partition
type Partition struct {
	cluster.Partition
	registry   *Registry
	ID         PartitionID
	services   map[ServiceId]Service
	servicesMu sync.RWMutex
	replicas   map[ServiceId]Replica
	replicasMu sync.RWMutex
}

func (p *Partition) GetService(ctx context.Context, serviceID ServiceId, timestamp *meta.Timestamp) (Service, error) {
	p.servicesMu.RLock()
	service, ok := p.services[serviceID]
	p.servicesMu.RUnlock()
	if ok {
		return service, nil
	}

	p.servicesMu.Lock()
	defer p.servicesMu.Unlock()
	service, ok = p.services[serviceID]
	if ok {
		return service, nil
	}

	f, err := p.registry.GetServiceFunc(serviceID.Type)
	if err != nil {
		return nil, err
	}
	service, err = f(ctx, serviceID, p, getClockFromTimestamp(*timestamp), getReplicationFactorFromContext(ctx))
	if err != nil {
		return nil, err
	}
	p.services[serviceID] = service
	return service, nil
}

func (p *Partition) RegisterReplica(replica Replica) error {
	p.replicasMu.Lock()
	defer p.replicasMu.Unlock()
	if _, ok := p.replicas[replica.ID()]; ok {
		return errors.NewAlreadyExists("replica '%s' already exists", replica.ID())
	}
	p.replicas[replica.ID()] = replica
	return nil
}

func (p *Partition) getReplica(ctx context.Context, serviceID ServiceId, timestamp *meta.Timestamp) (Replica, error) {
	p.replicasMu.RLock()
	replica, ok := p.replicas[serviceID]
	p.replicasMu.RUnlock()
	if !ok {
		_, err := p.GetService(ctx, serviceID, timestamp)
		if err != nil {
			return nil, err
		}
		p.replicasMu.RLock()
		replica, ok = p.replicas[serviceID]
		p.replicasMu.RUnlock()
		if !ok {
			return nil, errors.NewNotFound("replica '%s' not found", serviceID)
		}
	}
	if replica.ID().Type != serviceID.Type {
		return nil, errors.NewConflict("replica '%s' already exists with a different type '%s'", serviceID, replica.ID().Type)
	}
	return replica, nil
}

func (p *Partition) deleteReplica(serviceID ServiceId) error {
	p.servicesMu.Lock()
	defer p.servicesMu.Unlock()
	_, ok := p.services[serviceID]
	if !ok {
		return errors.NewNotFound("service '%s' not found", serviceID)
	}
	delete(p.services, serviceID)
	return nil
}
