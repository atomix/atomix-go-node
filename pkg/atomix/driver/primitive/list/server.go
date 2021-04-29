package list

import (
	"context"
	list "github.com/atomix/atomix-api/go/atomix/primitive/list"
	"github.com/atomix/atomix-go-framework/pkg/atomix/driver/env"
	"github.com/atomix/atomix-go-framework/pkg/atomix/logging"
)

// NewProxyServer creates a new ProxyServer
func NewProxyServer(registry *ProxyRegistry, env env.DriverEnv) list.ListServiceServer {
	return &ProxyServer{
		registry: registry,
		env:      env,
		log:      logging.GetLogger("atomix", "list"),
	}
}

type ProxyServer struct {
	registry *ProxyRegistry
	env      env.DriverEnv
	log      logging.Logger
}

func (s *ProxyServer) Size(ctx context.Context, request *list.SizeRequest) (*list.SizeResponse, error) {
	if request.Headers.PrimitiveID.Namespace == "" {
		request.Headers.PrimitiveID.Namespace = s.env.Namespace
	}
	proxy, err := s.registry.GetProxy(request.Headers.PrimitiveID)
	if err != nil {
		s.log.Warnf("SizeRequest %+v failed: %v", request, err)
		return nil, err
	}
	return proxy.Size(ctx, request)
}

func (s *ProxyServer) Append(ctx context.Context, request *list.AppendRequest) (*list.AppendResponse, error) {
	if request.Headers.PrimitiveID.Namespace == "" {
		request.Headers.PrimitiveID.Namespace = s.env.Namespace
	}
	proxy, err := s.registry.GetProxy(request.Headers.PrimitiveID)
	if err != nil {
		s.log.Warnf("AppendRequest %+v failed: %v", request, err)
		return nil, err
	}
	return proxy.Append(ctx, request)
}

func (s *ProxyServer) Insert(ctx context.Context, request *list.InsertRequest) (*list.InsertResponse, error) {
	if request.Headers.PrimitiveID.Namespace == "" {
		request.Headers.PrimitiveID.Namespace = s.env.Namespace
	}
	proxy, err := s.registry.GetProxy(request.Headers.PrimitiveID)
	if err != nil {
		s.log.Warnf("InsertRequest %+v failed: %v", request, err)
		return nil, err
	}
	return proxy.Insert(ctx, request)
}

func (s *ProxyServer) Get(ctx context.Context, request *list.GetRequest) (*list.GetResponse, error) {
	if request.Headers.PrimitiveID.Namespace == "" {
		request.Headers.PrimitiveID.Namespace = s.env.Namespace
	}
	proxy, err := s.registry.GetProxy(request.Headers.PrimitiveID)
	if err != nil {
		s.log.Warnf("GetRequest %+v failed: %v", request, err)
		return nil, err
	}
	return proxy.Get(ctx, request)
}

func (s *ProxyServer) Set(ctx context.Context, request *list.SetRequest) (*list.SetResponse, error) {
	if request.Headers.PrimitiveID.Namespace == "" {
		request.Headers.PrimitiveID.Namespace = s.env.Namespace
	}
	proxy, err := s.registry.GetProxy(request.Headers.PrimitiveID)
	if err != nil {
		s.log.Warnf("SetRequest %+v failed: %v", request, err)
		return nil, err
	}
	return proxy.Set(ctx, request)
}

func (s *ProxyServer) Remove(ctx context.Context, request *list.RemoveRequest) (*list.RemoveResponse, error) {
	if request.Headers.PrimitiveID.Namespace == "" {
		request.Headers.PrimitiveID.Namespace = s.env.Namespace
	}
	proxy, err := s.registry.GetProxy(request.Headers.PrimitiveID)
	if err != nil {
		s.log.Warnf("RemoveRequest %+v failed: %v", request, err)
		return nil, err
	}
	return proxy.Remove(ctx, request)
}

func (s *ProxyServer) Clear(ctx context.Context, request *list.ClearRequest) (*list.ClearResponse, error) {
	if request.Headers.PrimitiveID.Namespace == "" {
		request.Headers.PrimitiveID.Namespace = s.env.Namespace
	}
	proxy, err := s.registry.GetProxy(request.Headers.PrimitiveID)
	if err != nil {
		s.log.Warnf("ClearRequest %+v failed: %v", request, err)
		return nil, err
	}
	return proxy.Clear(ctx, request)
}

func (s *ProxyServer) Events(request *list.EventsRequest, srv list.ListService_EventsServer) error {
	if request.Headers.PrimitiveID.Namespace == "" {
		request.Headers.PrimitiveID.Namespace = s.env.Namespace
	}
	proxy, err := s.registry.GetProxy(request.Headers.PrimitiveID)
	if err != nil {
		s.log.Warnf("EventsRequest %+v failed: %v", request, err)
		return err
	}
	return proxy.Events(request, srv)
}

func (s *ProxyServer) Elements(request *list.ElementsRequest, srv list.ListService_ElementsServer) error {
	if request.Headers.PrimitiveID.Namespace == "" {
		request.Headers.PrimitiveID.Namespace = s.env.Namespace
	}
	proxy, err := s.registry.GetProxy(request.Headers.PrimitiveID)
	if err != nil {
		s.log.Warnf("ElementsRequest %+v failed: %v", request, err)
		return err
	}
	return proxy.Elements(request, srv)
}
