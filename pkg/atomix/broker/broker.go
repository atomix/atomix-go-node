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

package broker

import (
	brokerapi "github.com/atomix/atomix-api/go/atomix/management/broker"
	"github.com/atomix/atomix-go-framework/pkg/atomix/cluster"
	"github.com/atomix/atomix-go-framework/pkg/atomix/logging"
	"github.com/atomix/atomix-go-framework/pkg/atomix/server"
	"google.golang.org/grpc"
)

var log = logging.GetLogger("atomix", "broker")

// NewBroker creates a new broker node
func NewBroker(cluster cluster.Cluster, opts ...Option) *Broker {
	options := applyOptions(opts...)
	return &Broker{
		Server: server.NewServer(cluster),
		namespace: options.namespace,
	}
}

// Broker is a broker node
type Broker struct {
	*server.Server
	namespace string
}

// Start starts the node
func (n *Broker) Start() error {
	server := NewServer(newPrimitiveRegistry(n.namespace))
	n.Server.RegisterService(func(s *grpc.Server) {
		brokerapi.RegisterBrokerServer(s, server)
	})
	if err := n.Server.Start(); err != nil {
		return err
	}
	return nil
}

// Stop stops the node
func (n *Broker) Stop() error {
	if err := n.Server.Stop(); err != nil {
		return err
	}
	return nil
}
