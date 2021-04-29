package _map

import (
	"context"
	_map "github.com/atomix/atomix-api/go/atomix/primitive/map"
	"github.com/atomix/atomix-go-framework/pkg/atomix/driver/proxy/gossip"
	"github.com/atomix/atomix-go-framework/pkg/atomix/errors"
	"github.com/atomix/atomix-go-framework/pkg/atomix/logging"
	async "github.com/atomix/atomix-go-framework/pkg/atomix/util/async"

	io "io"
	sync "sync"
)

// NewProxyServer creates a new ProxyServer
func NewProxyServer(client *gossip.Client) _map.MapServiceServer {
	return &ProxyServer{
		Client: client,
		log:    logging.GetLogger("atomix", "map"),
	}
}

type ProxyServer struct {
	*gossip.Client
	log logging.Logger
}

func (s *ProxyServer) Size(ctx context.Context, request *_map.SizeRequest) (*_map.SizeResponse, error) {
	s.log.Debugf("Received SizeRequest %+v", request)
	partitions := s.Partitions()
	responses, err := async.ExecuteAsync(len(partitions), func(i int) (interface{}, error) {
		var prequest *_map.SizeRequest
		*prequest = *request
		partition := partitions[i]
		conn, err := partition.Connect()
		if err != nil {
			return nil, err
		}
		client := _map.NewMapServiceClient(conn)
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

	response := &_map.SizeResponse{}
	s.AddResponseHeaders(&response.Headers)
	for _, r := range responses {
		response.Size_ += r.(*_map.SizeResponse).Size_
	}
	s.log.Debugf("Sending SizeResponse %+v", response)
	return response, nil
}

func (s *ProxyServer) Put(ctx context.Context, request *_map.PutRequest) (*_map.PutResponse, error) {
	s.log.Debugf("Received PutRequest %+v", request)
	partitionKey := request.Entry.Key.Key
	partition := s.PartitionBy([]byte(partitionKey))

	conn, err := partition.Connect()
	if err != nil {
		return nil, errors.Proto(err)
	}

	client := _map.NewMapServiceClient(conn)
	partition.AddRequestHeaders(&request.Headers)
	response, err := client.Put(ctx, request)
	if err != nil {
		s.log.Errorf("Request PutRequest failed: %v", err)
		return nil, errors.Proto(err)
	}
	partition.AddResponseHeaders(&response.Headers)
	s.log.Debugf("Sending PutResponse %+v", response)
	return response, nil
}

func (s *ProxyServer) Get(ctx context.Context, request *_map.GetRequest) (*_map.GetResponse, error) {
	s.log.Debugf("Received GetRequest %+v", request)
	partitionKey := request.Key
	partition := s.PartitionBy([]byte(partitionKey))

	conn, err := partition.Connect()
	if err != nil {
		return nil, errors.Proto(err)
	}

	client := _map.NewMapServiceClient(conn)
	partition.AddRequestHeaders(&request.Headers)
	response, err := client.Get(ctx, request)
	if err != nil {
		s.log.Errorf("Request GetRequest failed: %v", err)
		return nil, errors.Proto(err)
	}
	partition.AddResponseHeaders(&response.Headers)
	s.log.Debugf("Sending GetResponse %+v", response)
	return response, nil
}

func (s *ProxyServer) Remove(ctx context.Context, request *_map.RemoveRequest) (*_map.RemoveResponse, error) {
	s.log.Debugf("Received RemoveRequest %+v", request)
	partitionKey := request.Key.Key
	partition := s.PartitionBy([]byte(partitionKey))

	conn, err := partition.Connect()
	if err != nil {
		return nil, errors.Proto(err)
	}

	client := _map.NewMapServiceClient(conn)
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

func (s *ProxyServer) Clear(ctx context.Context, request *_map.ClearRequest) (*_map.ClearResponse, error) {
	s.log.Debugf("Received ClearRequest %+v", request)
	partitions := s.Partitions()
	err := async.IterAsync(len(partitions), func(i int) error {
		var prequest *_map.ClearRequest
		*prequest = *request
		partition := partitions[i]
		conn, err := partition.Connect()
		if err != nil {
			return err
		}
		client := _map.NewMapServiceClient(conn)
		partition.AddRequestHeaders(&prequest.Headers)
		_, err = client.Clear(ctx, prequest)
		return err
	})
	if err != nil {
		s.log.Errorf("Request ClearRequest failed: %v", err)
		return nil, errors.Proto(err)
	}

	response := &_map.ClearResponse{}
	s.AddResponseHeaders(&response.Headers)
	s.log.Debugf("Sending ClearResponse %+v", response)
	return response, nil
}

func (s *ProxyServer) Events(request *_map.EventsRequest, srv _map.MapService_EventsServer) error {
	s.log.Debugf("Received EventsRequest %+v", request)
	partitions := s.Partitions()
	wg := &sync.WaitGroup{}
	responseCh := make(chan *_map.EventsResponse)
	errCh := make(chan error)
	err := async.IterAsync(len(partitions), func(i int) error {
		var prequest *_map.EventsRequest
		*prequest = *request
		partition := partitions[i]
		conn, err := partition.Connect()
		if err != nil {
			s.log.Errorf("Request EventsRequest failed: %v", err)
			return err
		}
		client := _map.NewMapServiceClient(conn)
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

func (s *ProxyServer) Entries(request *_map.EntriesRequest, srv _map.MapService_EntriesServer) error {
	s.log.Debugf("Received EntriesRequest %+v", request)
	partitions := s.Partitions()
	wg := &sync.WaitGroup{}
	responseCh := make(chan *_map.EntriesResponse)
	errCh := make(chan error)
	err := async.IterAsync(len(partitions), func(i int) error {
		var prequest *_map.EntriesRequest
		*prequest = *request
		partition := partitions[i]
		conn, err := partition.Connect()
		if err != nil {
			s.log.Errorf("Request EntriesRequest failed: %v", err)
			return err
		}
		client := _map.NewMapServiceClient(conn)
		partition.AddRequestHeaders(&prequest.Headers)
		stream, err := client.Entries(srv.Context(), prequest)
		if err != nil {
			s.log.Errorf("Request EntriesRequest failed: %v", err)
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
		s.log.Errorf("Request EntriesRequest failed: %v", err)
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
				s.log.Debugf("Sending EntriesResponse %+v", response)
				err := srv.Send(response)
				if err != nil {
					s.log.Errorf("Response EntriesResponse failed: %v", err)
					return err
				}
			}
		case err := <-errCh:
			if err != nil {
				s.log.Errorf("Request EntriesRequest failed: %v", err)
			}
			s.log.Debugf("Finished EntriesRequest %+v", request)
			return err
		}
	}
}
