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

package log

import (
	driverapi "github.com/atomix/atomix-api/go/atomix/management/driver"
	logapi "github.com/atomix/atomix-api/go/atomix/primitive/log"
	"github.com/atomix/atomix-go-framework/pkg/atomix/driver/primitive"
	logdriver "github.com/atomix/atomix-go-framework/pkg/atomix/driver/primitive/log"
	logro "github.com/atomix/atomix-go-framework/pkg/atomix/driver/proxy/ro/log"
	"github.com/atomix/atomix-go-framework/pkg/atomix/driver/proxy/rsm"
	"github.com/gogo/protobuf/jsonpb"
	"google.golang.org/grpc"
)

func Register(protocol *rsm.Protocol) {
	protocol.Primitives().RegisterPrimitiveType(newLogType(protocol))
}

func newLogType(protocol *rsm.Protocol) primitive.PrimitiveType {
	return &logType{
		protocol: protocol,
		registry: logdriver.NewProxyRegistry(),
	}
}

type logType struct {
	protocol *rsm.Protocol
	registry *logdriver.ProxyRegistry
}

func (p *logType) Name() string {
	return Type
}

func (p *logType) RegisterServer(s *grpc.Server) {
	logapi.RegisterLogServiceServer(s, logdriver.NewProxyServer(p.registry, p.protocol.Env))
}

func (p *logType) AddProxy(id driverapi.ProxyId, options driverapi.ProxyOptions) error {
	config := rsm.RSMConfig{}
	if options.Config != nil {
		if err := jsonpb.UnmarshalString(string(options.Config), &config); err != nil {
			return err
		}
	}
	server := NewProxyServer(p.protocol.Client, config.ReadSync)
	if !options.Write {
		server = logro.NewProxyServer(server)
	}
	return p.registry.AddProxy(id.PrimitiveId, server)
}

func (p *logType) RemoveProxy(id driverapi.ProxyId) error {
	return p.registry.RemoveProxy(id.PrimitiveId)
}

var _ primitive.PrimitiveType = &logType{}
