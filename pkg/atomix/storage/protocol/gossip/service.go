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
	"github.com/atomix/atomix-go-framework/pkg/atomix/time"
)

// ServiceType is a gossip service type name
type ServiceType string

// Service is a gossip service
type Service interface{}

// Replica is a service replica interface
type Replica interface {
	ID() ServiceId
	Clock() time.Clock
	Read(ctx context.Context, key string) (*Object, error)
	ReadAll(ctx context.Context, ch chan<- Object) error
	Update(ctx context.Context, object *Object) error
}
