

package set

import (
	"context"
	"github.com/atomix/go-framework/pkg/atomix/logging"
	set "github.com/atomix/api/go/atomix/primitive/set"
)

// NewProxyServer creates a new ProxyServer
func NewProxyServer(registry *ProxyRegistry) set.SetServiceServer {
	return &ProxyServer{
		registry: registry,
		log:      logging.GetLogger("atomix", "set"),
	}
}
type ProxyServer struct {
	registry *ProxyRegistry
	log      logging.Logger
}

func (s *ProxyServer) Size(ctx context.Context, request *set.SizeRequest) (*set.SizeResponse, error) {
	proxy, err := s.registry.GetProxy(request.Headers.PrimitiveID)
	if err != nil {
	    s.log.Warnf("SizeRequest %+v failed: %v", request, err)
		return nil, err
	}
	return proxy.Size(ctx, request)
}


func (s *ProxyServer) Contains(ctx context.Context, request *set.ContainsRequest) (*set.ContainsResponse, error) {
	proxy, err := s.registry.GetProxy(request.Headers.PrimitiveID)
	if err != nil {
	    s.log.Warnf("ContainsRequest %+v failed: %v", request, err)
		return nil, err
	}
	return proxy.Contains(ctx, request)
}


func (s *ProxyServer) Add(ctx context.Context, request *set.AddRequest) (*set.AddResponse, error) {
	proxy, err := s.registry.GetProxy(request.Headers.PrimitiveID)
	if err != nil {
	    s.log.Warnf("AddRequest %+v failed: %v", request, err)
		return nil, err
	}
	return proxy.Add(ctx, request)
}


func (s *ProxyServer) Remove(ctx context.Context, request *set.RemoveRequest) (*set.RemoveResponse, error) {
	proxy, err := s.registry.GetProxy(request.Headers.PrimitiveID)
	if err != nil {
	    s.log.Warnf("RemoveRequest %+v failed: %v", request, err)
		return nil, err
	}
	return proxy.Remove(ctx, request)
}


func (s *ProxyServer) Clear(ctx context.Context, request *set.ClearRequest) (*set.ClearResponse, error) {
	proxy, err := s.registry.GetProxy(request.Headers.PrimitiveID)
	if err != nil {
	    s.log.Warnf("ClearRequest %+v failed: %v", request, err)
		return nil, err
	}
	return proxy.Clear(ctx, request)
}


func (s *ProxyServer) Events(request *set.EventsRequest, srv set.SetService_EventsServer) error {
	proxy, err := s.registry.GetProxy(request.Headers.PrimitiveID)
	if err != nil {
	    s.log.Warnf("EventsRequest %+v failed: %v", request, err)
		return err
	}
	return proxy.Events(request, srv)
}


func (s *ProxyServer) Elements(request *set.ElementsRequest, srv set.SetService_ElementsServer) error {
	proxy, err := s.registry.GetProxy(request.Headers.PrimitiveID)
	if err != nil {
	    s.log.Warnf("ElementsRequest %+v failed: %v", request, err)
		return err
	}
	return proxy.Elements(request, srv)
}
