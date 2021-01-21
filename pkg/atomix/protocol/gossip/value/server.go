package value

import (
	"context"
	value "github.com/atomix/api/go/atomix/primitive/value"
	"github.com/atomix/go-framework/pkg/atomix/errors"
	"github.com/atomix/go-framework/pkg/atomix/logging"
	"github.com/atomix/go-framework/pkg/atomix/protocol/gossip"
	"github.com/atomix/go-framework/pkg/atomix/util/async"
	"google.golang.org/grpc"
	"sync"
)

// RegisterServer registers the primitive on the given node
func RegisterServer(node *gossip.Node) {
	node.RegisterServer(func(server *grpc.Server, manager *gossip.Manager) {
		value.RegisterValueServiceServer(server, newServer(manager))
	})
	node.RegisterServer(registerServerFunc)
}

var registerServerFunc gossip.RegisterServerFunc

func newServer(manager *gossip.Manager) value.ValueServiceServer {
	return &Server{
		manager: manager,
		log:     logging.GetLogger("atomix", "protocol", "gossip", "value"),
	}
}

type Server struct {
	manager *gossip.Manager
	log     logging.Logger
}

func (s *Server) Set(ctx context.Context, request *value.SetRequest) (*value.SetResponse, error) {
	s.log.Debugf("Received SetRequest %+v", request)
	partition, err := s.manager.PartitionFrom(ctx)
	if err != nil {
		s.log.Errorf("Request SetRequest %+v failed: %v", request, err)
		return nil, err
	}

	service, err := partition.ServiceFrom(ctx)
	if err != nil {
		s.log.Errorf("Request SetRequest %+v failed: %v", request, err)
		return nil, errors.Proto(err)
	}

	response, err := service.(Service).Set(ctx, request)
	if err != nil {
		s.log.Errorf("Request SetRequest %+v failed: %v", request, err)
		return nil, errors.Proto(err)
	}
	s.log.Debugf("Sending SetResponse %+v", response)
	return response, nil
}

func (s *Server) Get(ctx context.Context, request *value.GetRequest) (*value.GetResponse, error) {
	s.log.Debugf("Received GetRequest %+v", request)
	partition, err := s.manager.PartitionFrom(ctx)
	if err != nil {
		s.log.Errorf("Request GetRequest %+v failed: %v", request, err)
		return nil, err
	}

	service, err := partition.ServiceFrom(ctx)
	if err != nil {
		s.log.Errorf("Request GetRequest %+v failed: %v", request, err)
		return nil, errors.Proto(err)
	}

	response, err := service.(Service).Get(ctx, request)
	if err != nil {
		s.log.Errorf("Request GetRequest %+v failed: %v", request, err)
		return nil, errors.Proto(err)
	}
	s.log.Debugf("Sending GetResponse %+v", response)
	return response, nil
}

func (s *Server) Events(request *value.EventsRequest, srv value.ValueService_EventsServer) error {
	s.log.Debugf("Received EventsRequest %+v", request)

	partitions, err := s.manager.PartitionsFrom(srv.Context())
	if err != nil {
		s.log.Errorf("Request EventsRequest %+v failed: %v", request, err)
		return errors.Proto(err)
	}

	responseCh := make(chan value.EventsResponse)
	wg := &sync.WaitGroup{}
	wg.Add(len(partitions))
	err = async.IterAsync(len(partitions), func(i int) error {
		partition := partitions[i]
		service, err := partition.ServiceFrom(srv.Context())
		if err != nil {
			return err
		}

		partitionCh := make(chan value.EventsResponse)
		errCh := make(chan error)
		go func() {
			err := service.(Service).Events(srv.Context(), request, partitionCh)
			if err != nil {
				errCh <- err
			}
		}()

		defer wg.Done()
		for {
			select {
			case response, ok := <-partitionCh:
				if ok {
					responseCh <- response
				} else {
					return nil
				}
			case err := <-errCh:
				return err
			}
		}
	})
	if err != nil {
		s.log.Errorf("Request EventsRequest %+v failed: %v", request, err)
		return errors.Proto(err)
	}

	go func() {
		wg.Wait()
		close(responseCh)
	}()

	for {
		select {
		case response, ok := <-responseCh:
			if ok {
				s.log.Debugf("Sending EventsResponse %v", response)
				err = srv.Send(&response)
				if err != nil {
					s.log.Errorf("Request EventsRequest %+v failed: %v", request, err)
					return errors.Proto(err)
				}
			} else {
				s.log.Debugf("Finished EventsRequest %+v", request)
				return nil
			}
		case <-srv.Context().Done():
			s.log.Debugf("Finished EventsRequest %+v", request)
			return nil
		}
	}
}