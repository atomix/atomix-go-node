package set

import (
	"context"
	set "github.com/atomix/atomix-api/go/atomix/primitive/set"
	"github.com/atomix/atomix-go-framework/pkg/atomix/driver/proxy/gossip"
	"github.com/atomix/atomix-go-framework/pkg/atomix/errors"
	"github.com/atomix/atomix-go-framework/pkg/atomix/logging"
	async "github.com/atomix/atomix-go-framework/pkg/atomix/util/async"

	io "io"
	sync "sync"
)

// NewProxyServer creates a new ProxyServer
func NewProxyServer(client *gossip.Client) set.SetServiceServer {
	return &ProxyServer{
		Client: client,
		log:    logging.GetLogger("atomix", "set"),
	}
}

type ProxyServer struct {
	*gossip.Client
	log logging.Logger
}

func (s *ProxyServer) Size(ctx context.Context, request *set.SizeRequest) (*set.SizeResponse, error) {
	s.log.Debugf("Received SizeRequest %+v", request)
	partitions := s.Partitions()
	responses, err := async.ExecuteAsync(len(partitions), func(i int) (interface{}, error) {
		var prequest *set.SizeRequest
		*prequest = *request
		partition := partitions[i]
		conn, err := partition.Connect()
		if err != nil {
			return nil, err
		}
		client := set.NewSetServiceClient(conn)
		partition.AddRequestHeaders(&prequest.Headers)
		presponse, err := client.Size(ctx, prequest)
		if err != nil {
			return nil, err
		}
		partition.AddResponseHeaders(&presponse.Headers)
		return presponse, nil
	})
	if err != nil {
		s.log.Errorf("Request SizeRequest failed: %v", err)
		return nil, errors.Proto(err)
	}

	response := &set.SizeResponse{}
	s.AddResponseHeaders(&response.Headers)
	for _, r := range responses {
		response.Size_ += r.(*set.SizeResponse).Size_
	}
	s.log.Debugf("Sending SizeResponse %+v", response)
	return response, nil
}

func (s *ProxyServer) Contains(ctx context.Context, request *set.ContainsRequest) (*set.ContainsResponse, error) {
	s.log.Debugf("Received ContainsRequest %+v", request)
	partitionKey := request.Element.Value
	partition := s.PartitionBy([]byte(partitionKey))

	conn, err := partition.Connect()
	if err != nil {
		return nil, errors.Proto(err)
	}

	client := set.NewSetServiceClient(conn)
	partition.AddRequestHeaders(&request.Headers)
	response, err := client.Contains(ctx, request)
	if err != nil {
		s.log.Errorf("Request ContainsRequest failed: %v", err)
		return nil, errors.Proto(err)
	}
	partition.AddResponseHeaders(&response.Headers)
	s.log.Debugf("Sending ContainsResponse %+v", response)
	return response, nil
}

func (s *ProxyServer) Add(ctx context.Context, request *set.AddRequest) (*set.AddResponse, error) {
	s.log.Debugf("Received AddRequest %+v", request)
	partitionKey := request.Element.Value
	partition := s.PartitionBy([]byte(partitionKey))

	conn, err := partition.Connect()
	if err != nil {
		return nil, errors.Proto(err)
	}

	client := set.NewSetServiceClient(conn)
	partition.AddRequestHeaders(&request.Headers)
	response, err := client.Add(ctx, request)
	if err != nil {
		s.log.Errorf("Request AddRequest failed: %v", err)
		return nil, errors.Proto(err)
	}
	partition.AddResponseHeaders(&response.Headers)
	s.log.Debugf("Sending AddResponse %+v", response)
	return response, nil
}

func (s *ProxyServer) Remove(ctx context.Context, request *set.RemoveRequest) (*set.RemoveResponse, error) {
	s.log.Debugf("Received RemoveRequest %+v", request)
	partitionKey := request.Element.Value
	partition := s.PartitionBy([]byte(partitionKey))

	conn, err := partition.Connect()
	if err != nil {
		return nil, errors.Proto(err)
	}

	client := set.NewSetServiceClient(conn)
	partition.AddRequestHeaders(&request.Headers)
	response, err := client.Remove(ctx, request)
	if err != nil {
		s.log.Errorf("Request RemoveRequest failed: %v", err)
		return nil, errors.Proto(err)
	}
	partition.AddResponseHeaders(&response.Headers)
	s.log.Debugf("Sending RemoveResponse %+v", response)
	return response, nil
}

func (s *ProxyServer) Clear(ctx context.Context, request *set.ClearRequest) (*set.ClearResponse, error) {
	s.log.Debugf("Received ClearRequest %+v", request)
	partitions := s.Partitions()
	err := async.IterAsync(len(partitions), func(i int) error {
		var prequest *set.ClearRequest
		*prequest = *request
		partition := partitions[i]
		conn, err := partition.Connect()
		if err != nil {
			return err
		}
		client := set.NewSetServiceClient(conn)
		partition.AddRequestHeaders(&prequest.Headers)
		_, err = client.Clear(ctx, prequest)
		return err
	})
	if err != nil {
		s.log.Errorf("Request ClearRequest failed: %v", err)
		return nil, errors.Proto(err)
	}

	response := &set.ClearResponse{}
	s.AddResponseHeaders(&response.Headers)
	s.log.Debugf("Sending ClearResponse %+v", response)
	return response, nil
}

func (s *ProxyServer) Events(request *set.EventsRequest, srv set.SetService_EventsServer) error {
	s.log.Debugf("Received EventsRequest %+v", request)
	partitions := s.Partitions()
	wg := &sync.WaitGroup{}
	responseCh := make(chan *set.EventsResponse)
	errCh := make(chan error)
	err := async.IterAsync(len(partitions), func(i int) error {
		var prequest *set.EventsRequest
		*prequest = *request
		partition := partitions[i]
		conn, err := partition.Connect()
		if err != nil {
			s.log.Errorf("Request EventsRequest failed: %v", err)
			return err
		}
		client := set.NewSetServiceClient(conn)
		partition.AddRequestHeaders(&prequest.Headers)
		stream, err := client.Events(srv.Context(), prequest)
		if err != nil {
			s.log.Errorf("Request EventsRequest failed: %v", err)
			return err
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				presponse, err := stream.Recv()
				if err == io.EOF {
					return
				} else if err != nil {
					errCh <- err
				} else {
					partition.AddResponseHeaders(&presponse.Headers)
					responseCh <- presponse
				}
			}
		}()
		return nil
	})
	if err != nil {
		s.log.Errorf("Request EventsRequest failed: %v", err)
		return errors.Proto(err)
	}

	go func() {
		wg.Wait()
		close(responseCh)
		close(errCh)
	}()

	for {
		select {
		case response, ok := <-responseCh:
			if ok {
				s.AddResponseHeaders(&response.Headers)
				s.log.Debugf("Sending EventsResponse %+v", response)
				err := srv.Send(response)
				if err != nil {
					s.log.Errorf("Response EventsResponse failed: %v", err)
					return err
				}
			}
		case err := <-errCh:
			if err != nil {
				s.log.Errorf("Request EventsRequest failed: %v", err)
			}
			s.log.Debugf("Finished EventsRequest %+v", request)
			return err
		}
	}
}

func (s *ProxyServer) Elements(request *set.ElementsRequest, srv set.SetService_ElementsServer) error {
	s.log.Debugf("Received ElementsRequest %+v", request)
	partitions := s.Partitions()
	wg := &sync.WaitGroup{}
	responseCh := make(chan *set.ElementsResponse)
	errCh := make(chan error)
	err := async.IterAsync(len(partitions), func(i int) error {
		var prequest *set.ElementsRequest
		*prequest = *request
		partition := partitions[i]
		conn, err := partition.Connect()
		if err != nil {
			s.log.Errorf("Request ElementsRequest failed: %v", err)
			return err
		}
		client := set.NewSetServiceClient(conn)
		partition.AddRequestHeaders(&prequest.Headers)
		stream, err := client.Elements(srv.Context(), prequest)
		if err != nil {
			s.log.Errorf("Request ElementsRequest failed: %v", err)
			return err
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				presponse, err := stream.Recv()
				if err == io.EOF {
					return
				} else if err != nil {
					errCh <- err
				} else {
					partition.AddResponseHeaders(&presponse.Headers)
					responseCh <- presponse
				}
			}
		}()
		return nil
	})
	if err != nil {
		s.log.Errorf("Request ElementsRequest failed: %v", err)
		return errors.Proto(err)
	}

	go func() {
		wg.Wait()
		close(responseCh)
		close(errCh)
	}()

	for {
		select {
		case response, ok := <-responseCh:
			if ok {
				s.AddResponseHeaders(&response.Headers)
				s.log.Debugf("Sending ElementsResponse %+v", response)
				err := srv.Send(response)
				if err != nil {
					s.log.Errorf("Response ElementsResponse failed: %v", err)
					return err
				}
			}
		case err := <-errCh:
			if err != nil {
				s.log.Errorf("Request ElementsRequest failed: %v", err)
			}
			s.log.Debugf("Finished ElementsRequest %+v", request)
			return err
		}
	}
}
