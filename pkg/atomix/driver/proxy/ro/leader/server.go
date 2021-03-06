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

package leader

import (
	"context"
	leaderapi "github.com/atomix/atomix-api/go/atomix/primitive/leader"
	"github.com/atomix/atomix-go-framework/pkg/atomix/errors"
)

// NewProxyServer creates a new read-only leader server
func NewProxyServer(s leaderapi.LeaderLatchServiceServer) leaderapi.LeaderLatchServiceServer {
	return &ProxyServer{
		server: s,
	}
}

// ProxyServer is a read-only leader primitive server
type ProxyServer struct {
	server leaderapi.LeaderLatchServiceServer
}

func (r *ProxyServer) Latch(ctx context.Context, request *leaderapi.LatchRequest) (*leaderapi.LatchResponse, error) {
	return nil, errors.NewUnauthorized("Latch operation is not permitted")
}

func (r *ProxyServer) Get(ctx context.Context, request *leaderapi.GetRequest) (*leaderapi.GetResponse, error) {
	return r.server.Get(ctx, request)
}

func (r *ProxyServer) Events(request *leaderapi.EventsRequest, server leaderapi.LeaderLatchService_EventsServer) error {
	return r.server.Events(request, server)
}

var _ leaderapi.LeaderLatchServiceServer = &ProxyServer{}
