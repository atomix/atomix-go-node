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

package value

import (
	"github.com/atomix/atomix-go-node/pkg/atomix/service"
	"github.com/golang/protobuf/proto"
)

// RegisterValueService registers the set service in the given service registry
func RegisterValueService(registry *service.ServiceRegistry) {
	registry.Register("value", newValueService)
}

// newValueService returns a new ValueService
func newValueService(context service.Context) service.Service {
	service := &ValueService{
		SessionizedService: service.NewSessionizedService(context),
	}
	service.init()
	return service
}

// ValueService is a state machine for a list primitive
type ValueService struct {
	*service.SessionizedService
	value   []byte
	version uint64
}

// init initializes the list service
func (v *ValueService) init() {
	v.Executor.Register("set", v.Set)
	v.Executor.Register("get", v.Get)
	v.Executor.Register("events", v.Events)
}

// Backup backs up the value service
func (v *ValueService) Backup() ([]byte, error) {
	snapshot := &ValueSnapshot{
		Value:   v.value,
		Version: v.version,
	}
	return proto.Marshal(snapshot)
}

// Restore restores the value service
func (v *ValueService) Restore(bytes []byte) error {
	snapshot := &ValueSnapshot{}
	if err := proto.Unmarshal(bytes, snapshot); err != nil {
		return err
	}
	v.value = snapshot.Value
	v.version = snapshot.Version
	return nil
}

func (v *ValueService) Set(bytes []byte, ch chan<- service.Result) {
	defer close(ch)

	request := &SetRequest{}
	if err := proto.Unmarshal(bytes, request); err != nil {
		ch <- v.NewFailure(err)
		return
	}

	if request.ExpectVersion > 0 && request.ExpectVersion != v.version {
		ch <- v.NewResult(proto.Marshal(&SetResponse{
			Version:   v.version,
			Succeeded: false,
		}))
	} else if request.ExpectValue != nil && len(request.ExpectValue) > 0 && (v.value == nil || !slicesEqual(v.value, request.ExpectValue)) {
		ch <- v.NewResult(proto.Marshal(&SetResponse{
			Version:   v.version,
			Succeeded: false,
		}))
	} else {
		prevValue := v.value
		prevVersion := v.version
		v.value = request.Value
		v.version++

		v.sendEvent(&ListenResponse{
			Type:            ListenResponse_UPDATED,
			PreviousValue:   prevValue,
			PreviousVersion: prevVersion,
			NewValue:        v.value,
			NewVersion:      v.version,
		})

		ch <- v.NewResult(proto.Marshal(&SetResponse{
			Version:   v.version,
			Succeeded: true,
		}))
	}
}

func (v *ValueService) Get(bytes []byte, ch chan<- service.Result) {
	defer close(ch)

	request := &GetRequest{}
	if err := proto.Unmarshal(bytes, request); err != nil {
		ch <- v.NewFailure(err)
		return
	}

	ch <- v.NewResult(proto.Marshal(&GetResponse{
		Value:   v.value,
		Version: v.version,
	}))
}

func (v *ValueService) Events(bytes []byte, ch chan<- service.Result) {
	// Keep the stream open
}

func (v *ValueService) sendEvent(event *ListenResponse) {
	bytes, err := proto.Marshal(event)
	for _, session := range v.Sessions() {
		for _, ch := range session.ChannelsOf("events") {
			ch <- v.NewResult(bytes, err)
		}
	}
}

func slicesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for _, i := range a {
		for _, j := range b {
			if i != j {
				return false
			}
		}
	}
	return true
}