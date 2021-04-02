

package lock

import (
	"github.com/atomix/go-framework/pkg/atomix/storage/protocol/rsm"
	"github.com/atomix/go-framework/pkg/atomix/logging"
	"github.com/golang/protobuf/proto"
	lock "github.com/atomix/api/go/atomix/primitive/lock"
)

const Type = "Lock"

const (
    lockOp = "Lock"
    unlockOp = "Unlock"
    getLockOp = "GetLock"
)

var newServiceFunc rsm.NewServiceFunc

func registerServiceFunc(rsmf NewServiceFunc) {
	newServiceFunc = func(scheduler rsm.Scheduler, context rsm.ServiceContext) rsm.Service {
		service := &ServiceAdaptor{
			Service: rsm.NewService(scheduler, context),
			rsm:     rsmf(scheduler, context),
			log:     logging.GetLogger("atomix", "lock", "service"),
		}
		service.init()
		return service
	}
}

type NewServiceFunc func(scheduler rsm.Scheduler, context rsm.ServiceContext) Service

// RegisterService registers the election primitive service on the given node
func RegisterService(node *rsm.Node) {
	node.RegisterService(Type, newServiceFunc)
}

type ServiceAdaptor struct {
	rsm.Service
	rsm Service
	log logging.Logger
}

func (s *ServiceAdaptor) init() {
	s.RegisterStreamOperation(lockOp, s.lock)
	s.RegisterUnaryOperation(unlockOp, s.unlock)
	s.RegisterUnaryOperation(getLockOp, s.getLock)
}

func (s *ServiceAdaptor) SessionOpen(session rsm.Session) {
    if sessionOpen, ok := s.rsm.(rsm.SessionOpenService); ok {
        sessionOpen.SessionOpen(session)
    }
}

func (s *ServiceAdaptor) SessionExpired(session rsm.Session) {
    if sessionExpired, ok := s.rsm.(rsm.SessionExpiredService); ok {
        sessionExpired.SessionExpired(session)
    }
}

func (s *ServiceAdaptor) SessionClosed(session rsm.Session) {
    if sessionClosed, ok := s.rsm.(rsm.SessionClosedService); ok {
        sessionClosed.SessionClosed(session)
    }
}

func (s *ServiceAdaptor) lock(input []byte, stream rsm.Stream) (rsm.StreamCloser, error) {
    request := &lock.LockRequest{}
    err := proto.Unmarshal(input, request)
    if err != nil {
        s.log.Error(err)
        return nil, err
    }
    future, err := s.rsm.Lock(request)
    if err != nil {
        s.log.Error(err)
        return nil, err
    }
    future.setStream(stream)
    return nil, nil
}


func (s *ServiceAdaptor) unlock(input []byte) ([]byte, error) {
    request := &lock.UnlockRequest{}
	err := proto.Unmarshal(input, request)
	if err != nil {
	    s.log.Error(err)
		return nil, err
	}

	response, err := s.rsm.Unlock(request)
	if err !=  nil {
	    s.log.Error(err)
    	return nil, err
	}

	output, err := proto.Marshal(response)
	if err != nil {
	    s.log.Error(err)
		return nil, err
	}
	return output, nil
}


func (s *ServiceAdaptor) getLock(input []byte) ([]byte, error) {
    request := &lock.GetLockRequest{}
	err := proto.Unmarshal(input, request)
	if err != nil {
	    s.log.Error(err)
		return nil, err
	}

	response, err := s.rsm.GetLock(request)
	if err !=  nil {
	    s.log.Error(err)
    	return nil, err
	}

	output, err := proto.Marshal(response)
	if err != nil {
	    s.log.Error(err)
		return nil, err
	}
	return output, nil
}


var _ rsm.Service = &ServiceAdaptor{}