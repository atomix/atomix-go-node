package election

import (
	"context"
	election "github.com/atomix/atomix-api/go/atomix/primitive/election"
	"github.com/atomix/atomix-go-framework/pkg/atomix/driver/env"
	"github.com/atomix/atomix-go-framework/pkg/atomix/logging"
)

// NewProxyServer creates a new ProxyServer
func NewProxyServer(registry *ProxyRegistry, env env.DriverEnv) election.LeaderElectionServiceServer {
	return &ProxyServer{
		registry: registry,
		env:      env,
		log:      logging.GetLogger("atomix", "election"),
	}
}

type ProxyServer struct {
	registry *ProxyRegistry
	env      env.DriverEnv
	log      logging.Logger
}

func (s *ProxyServer) Enter(ctx context.Context, request *election.EnterRequest) (*election.EnterResponse, error) {
	if request.Headers.PrimitiveID.Namespace == "" {
		request.Headers.PrimitiveID.Namespace = s.env.Namespace
	}
	proxy, err := s.registry.GetProxy(request.Headers.PrimitiveID)
	if err != nil {
		s.log.Warnf("EnterRequest %+v failed: %v", request, err)
		return nil, err
	}
	return proxy.Enter(ctx, request)
}

func (s *ProxyServer) Withdraw(ctx context.Context, request *election.WithdrawRequest) (*election.WithdrawResponse, error) {
	if request.Headers.PrimitiveID.Namespace == "" {
		request.Headers.PrimitiveID.Namespace = s.env.Namespace
	}
	proxy, err := s.registry.GetProxy(request.Headers.PrimitiveID)
	if err != nil {
		s.log.Warnf("WithdrawRequest %+v failed: %v", request, err)
		return nil, err
	}
	return proxy.Withdraw(ctx, request)
}

func (s *ProxyServer) Anoint(ctx context.Context, request *election.AnointRequest) (*election.AnointResponse, error) {
	if request.Headers.PrimitiveID.Namespace == "" {
		request.Headers.PrimitiveID.Namespace = s.env.Namespace
	}
	proxy, err := s.registry.GetProxy(request.Headers.PrimitiveID)
	if err != nil {
		s.log.Warnf("AnointRequest %+v failed: %v", request, err)
		return nil, err
	}
	return proxy.Anoint(ctx, request)
}

func (s *ProxyServer) Promote(ctx context.Context, request *election.PromoteRequest) (*election.PromoteResponse, error) {
	if request.Headers.PrimitiveID.Namespace == "" {
		request.Headers.PrimitiveID.Namespace = s.env.Namespace
	}
	proxy, err := s.registry.GetProxy(request.Headers.PrimitiveID)
	if err != nil {
		s.log.Warnf("PromoteRequest %+v failed: %v", request, err)
		return nil, err
	}
	return proxy.Promote(ctx, request)
}

func (s *ProxyServer) Evict(ctx context.Context, request *election.EvictRequest) (*election.EvictResponse, error) {
	if request.Headers.PrimitiveID.Namespace == "" {
		request.Headers.PrimitiveID.Namespace = s.env.Namespace
	}
	proxy, err := s.registry.GetProxy(request.Headers.PrimitiveID)
	if err != nil {
		s.log.Warnf("EvictRequest %+v failed: %v", request, err)
		return nil, err
	}
	return proxy.Evict(ctx, request)
}

func (s *ProxyServer) GetTerm(ctx context.Context, request *election.GetTermRequest) (*election.GetTermResponse, error) {
	if request.Headers.PrimitiveID.Namespace == "" {
		request.Headers.PrimitiveID.Namespace = s.env.Namespace
	}
	proxy, err := s.registry.GetProxy(request.Headers.PrimitiveID)
	if err != nil {
		s.log.Warnf("GetTermRequest %+v failed: %v", request, err)
		return nil, err
	}
	return proxy.GetTerm(ctx, request)
}

func (s *ProxyServer) Events(request *election.EventsRequest, srv election.LeaderElectionService_EventsServer) error {
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
