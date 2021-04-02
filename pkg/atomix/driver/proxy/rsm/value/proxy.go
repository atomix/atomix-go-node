
package value

import (
	"context"
	"github.com/atomix/go-framework/pkg/atomix/driver/proxy/rsm"
	storage "github.com/atomix/go-framework/pkg/atomix/storage/protocol/rsm"
	"github.com/atomix/go-framework/pkg/atomix/errors"
	"github.com/atomix/go-framework/pkg/atomix/logging"
	"github.com/golang/protobuf/proto"
	value "github.com/atomix/api/go/atomix/primitive/value"
	streams "github.com/atomix/go-framework/pkg/atomix/stream"
)

const Type = "Value"

const (
    setOp = "Set"
    getOp = "Get"
    eventsOp = "Events"
)

// NewProxyServer creates a new ProxyServer
func NewProxyServer(client *rsm.Client) value.ValueServiceServer {
	return &ProxyServer{
		Client: client,
		log:    logging.GetLogger("atomix", "counter"),
	}
}

type ProxyServer struct {
	*rsm.Client
	log logging.Logger
}

func (s *ProxyServer) Set(ctx context.Context, request *value.SetRequest) (*value.SetResponse, error) {
	s.log.Debugf("Received SetRequest %+v", request)
	input, err := proto.Marshal(request)
	if err != nil {
        s.log.Errorf("Request SetRequest failed: %v", err)
	    return nil, errors.Proto(err)
	}
    partition := s.PartitionBy([]byte(request.Headers.PrimitiveID.String()))

	service := storage.ServiceId{
		Type:      Type,
		Namespace: request.Headers.PrimitiveID.Namespace,
		Name:      request.Headers.PrimitiveID.Name,
	}
	output, err := partition.DoCommand(ctx, service, setOp, input)
	if err != nil {
        s.log.Errorf("Request SetRequest failed: %v", err)
	    return nil, errors.Proto(err)
	}

	response := &value.SetResponse{}
	err = proto.Unmarshal(output, response)
	if err != nil {
        s.log.Errorf("Request SetRequest failed: %v", err)
	    return nil, errors.Proto(err)
	}
	s.log.Debugf("Sending SetResponse %+v", response)
	return response, nil
}


func (s *ProxyServer) Get(ctx context.Context, request *value.GetRequest) (*value.GetResponse, error) {
	s.log.Debugf("Received GetRequest %+v", request)
	input, err := proto.Marshal(request)
	if err != nil {
        s.log.Errorf("Request GetRequest failed: %v", err)
	    return nil, errors.Proto(err)
	}
    partition := s.PartitionBy([]byte(request.Headers.PrimitiveID.String()))

	service := storage.ServiceId{
		Type:      Type,
		Namespace: request.Headers.PrimitiveID.Namespace,
		Name:      request.Headers.PrimitiveID.Name,
	}
	output, err := partition.DoQuery(ctx, service, getOp, input)
	if err != nil {
        s.log.Errorf("Request GetRequest failed: %v", err)
	    return nil, errors.Proto(err)
	}

	response := &value.GetResponse{}
	err = proto.Unmarshal(output, response)
	if err != nil {
        s.log.Errorf("Request GetRequest failed: %v", err)
	    return nil, errors.Proto(err)
	}
	s.log.Debugf("Sending GetResponse %+v", response)
	return response, nil
}


func (s *ProxyServer) Events(request *value.EventsRequest, srv value.ValueService_EventsServer) error {
    s.log.Debugf("Received EventsRequest %+v", request)
	input, err := proto.Marshal(request)
	if err != nil {
        s.log.Errorf("Request EventsRequest failed: %v", err)
        return errors.Proto(err)
	}

	stream := streams.NewBufferedStream()
    partition := s.PartitionBy([]byte(request.Headers.PrimitiveID.String()))

	service := storage.ServiceId{
		Type:      Type,
		Namespace: request.Headers.PrimitiveID.Namespace,
		Name:      request.Headers.PrimitiveID.Name,
	}
	err = partition.DoCommandStream(srv.Context(), service, eventsOp, input, stream)
	if err != nil {
        s.log.Errorf("Request EventsRequest failed: %v", err)
	    return errors.Proto(err)
	}

	for {
		result, ok := stream.Receive()
		if !ok {
			break
		}

		if result.Failed() {
			s.log.Errorf("Request EventsRequest failed: %v", result.Error)
			return errors.Proto(result.Error)
		}

		response := &value.EventsResponse{}
        err = proto.Unmarshal(result.Value.([]byte), response)
        if err != nil {
            s.log.Errorf("Request EventsRequest failed: %v", err)
            return errors.Proto(err)
        }

		s.log.Debugf("Sending EventsResponse %+v", response)
		if err = srv.Send(response); err != nil {
            s.log.Errorf("Response EventsResponse failed: %v", err)
			return errors.Proto(err)
		}
	}
	s.log.Debugf("Finished EventsRequest %+v", request)
	return nil
}
