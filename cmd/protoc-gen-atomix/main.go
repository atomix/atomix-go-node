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

package main

import (
	"github.com/atomix/atomix-go-framework/codegen"
	"github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

func main() {
	pgs.Init(pgs.DebugMode()).
		RegisterModule(codegen.NewModule("driver", "primitive", map[string]string{
			"server.go":   "/etc/atomix/templates/driver/primitive/server.tpl",
			"registry.go": "/etc/atomix/templates/driver/primitive/registry.tpl",
		})).
		RegisterModule(codegen.NewModule("proxy", "rsm", map[string]string{
			"proxy.go": "/etc/atomix/templates/driver/proxy/rsm/proxy.tpl",
		})).
		RegisterModule(codegen.NewModule("storage", "rsm", map[string]string{
			"interface.go": "/etc/atomix/templates/storage/protocol/rsm/interface.tpl",
			"adapter.go":   "/etc/atomix/templates/storage/protocol/rsm/adapter.tpl",
		})).
		RegisterModule(codegen.NewModule("proxy", "gossip", map[string]string{
			"proxy.go": "/etc/atomix/templates/driver/proxy/gossip/proxy.tpl",
		})).
		RegisterModule(codegen.NewModule("storage", "gossip", map[string]string{
			"gossip.go":  "/etc/atomix/templates/storage/protocol/gossip/gossip.tpl",
			"server.go":  "/etc/atomix/templates/storage/protocol/gossip/server.tpl",
			"service.go": "/etc/atomix/templates/storage/protocol/gossip/service.tpl",
		})).
		RegisterPostProcessor(pgsgo.GoFmt()).
		Render()
}
