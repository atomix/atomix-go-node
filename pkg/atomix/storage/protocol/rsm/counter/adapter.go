// Code generated by atomix-go-framework. DO NOT EDIT.
package counter

import (
	counter "github.com/atomix/atomix-api/go/atomix/primitive/counter"
	"github.com/atomix/atomix-go-framework/pkg/atomix/errors"
	"github.com/atomix/atomix-go-framework/pkg/atomix/logging"
	"github.com/atomix/atomix-go-framework/pkg/atomix/storage/protocol/rsm"
	"github.com/atomix/atomix-go-framework/pkg/atomix/util"
	"github.com/golang/protobuf/proto"
	"io"
)

const Type = "Counter"

const (
	setOp       = "Set"
	getOp       = "Get"
	incrementOp = "Increment"
	decrementOp = "Decrement"
)

var newServiceFunc rsm.NewServiceFunc

func registerServiceFunc(rsmf NewServiceFunc) {
	newServiceFunc = func(scheduler rsm.Scheduler, context rsm.ServiceContext) rsm.Service {
		service := &ServiceAdaptor{
			Service: rsm.NewService(scheduler, context),
			rsm:     rsmf(newServiceContext(scheduler)),
			log:     logging.GetLogger("atomix", "counter", "service"),
		}
		service.init()
		return service
	}
}

type NewServiceFunc func(ServiceContext) Service

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
	s.RegisterUnaryOperation(setOp, s.set)
	s.RegisterUnaryOperation(getOp, s.get)
	s.RegisterUnaryOperation(incrementOp, s.increment)
	s.RegisterUnaryOperation(decrementOp, s.decrement)
}
func (s *ServiceAdaptor) SessionOpen(rsmSession rsm.Session) {
	s.rsm.Sessions().open(newSession(rsmSession))
}

func (s *ServiceAdaptor) SessionExpired(session rsm.Session) {
	s.rsm.Sessions().expire(SessionID(session.ID()))
}

func (s *ServiceAdaptor) SessionClosed(session rsm.Session) {
	s.rsm.Sessions().close(SessionID(session.ID()))
}
func (s *ServiceAdaptor) Backup(writer io.Writer) error {
	state, err := s.rsm.GetState()
	if err != nil {
		s.log.Error(err)
		return err
	}
	bytes, err := proto.Marshal(state)
	if err != nil {
		s.log.Error(err)
		return err
	}
	err = util.WriteBytes(writer, bytes)
	if err != nil {
		s.log.Error(err)
		return err
	}
	return nil
}

func (s *ServiceAdaptor) Restore(reader io.Reader) error {
	bytes, err := util.ReadBytes(reader)
	if err != nil {
		s.log.Error(err)
		return err
	}
	state := &CounterState{}
	err = proto.Unmarshal(bytes, state)
	if err != nil {
		s.log.Error(err)
		return err
	}
	err = s.rsm.SetState(state)
	if err != nil {
		s.log.Error(err)
		return err
	}
	return nil
}
func (s *ServiceAdaptor) set(input []byte, rsmSession rsm.Session) ([]byte, error) {
	request := &counter.SetRequest{}
	err := proto.Unmarshal(input, request)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}

	session, ok := s.rsm.Sessions().Get(SessionID(rsmSession.ID()))
	if !ok {
		err := errors.NewConflict("session %d not found", rsmSession.ID())
		s.log.Error(err)
		return nil, err
	}

	var response *counter.SetResponse
	proposal := newSetProposal(ProposalID(s.Index()), session, request, response)

	s.rsm.Proposals().Set().register(proposal)
	session.Proposals().Set().register(proposal)

	defer func() {
		session.Proposals().Set().unregister(proposal.ID())
		s.rsm.Proposals().Set().unregister(proposal.ID())
	}()

	err = s.rsm.Set(proposal)
	if err != nil {
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
func (s *ServiceAdaptor) get(input []byte, rsmSession rsm.Session) ([]byte, error) {
	request := &counter.GetRequest{}
	err := proto.Unmarshal(input, request)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}

	session, ok := s.rsm.Sessions().Get(SessionID(rsmSession.ID()))
	if !ok {
		err := errors.NewConflict("session %d not found", rsmSession.ID())
		s.log.Error(err)
		return nil, err
	}

	var response *counter.GetResponse
	proposal := newGetProposal(ProposalID(s.Index()), session, request, response)

	s.rsm.Proposals().Get().register(proposal)
	session.Proposals().Get().register(proposal)

	defer func() {
		session.Proposals().Get().unregister(proposal.ID())
		s.rsm.Proposals().Get().unregister(proposal.ID())
	}()

	err = s.rsm.Get(proposal)
	if err != nil {
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
func (s *ServiceAdaptor) increment(input []byte, rsmSession rsm.Session) ([]byte, error) {
	request := &counter.IncrementRequest{}
	err := proto.Unmarshal(input, request)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}

	session, ok := s.rsm.Sessions().Get(SessionID(rsmSession.ID()))
	if !ok {
		err := errors.NewConflict("session %d not found", rsmSession.ID())
		s.log.Error(err)
		return nil, err
	}

	var response *counter.IncrementResponse
	proposal := newIncrementProposal(ProposalID(s.Index()), session, request, response)

	s.rsm.Proposals().Increment().register(proposal)
	session.Proposals().Increment().register(proposal)

	defer func() {
		session.Proposals().Increment().unregister(proposal.ID())
		s.rsm.Proposals().Increment().unregister(proposal.ID())
	}()

	err = s.rsm.Increment(proposal)
	if err != nil {
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
func (s *ServiceAdaptor) decrement(input []byte, rsmSession rsm.Session) ([]byte, error) {
	request := &counter.DecrementRequest{}
	err := proto.Unmarshal(input, request)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}

	session, ok := s.rsm.Sessions().Get(SessionID(rsmSession.ID()))
	if !ok {
		err := errors.NewConflict("session %d not found", rsmSession.ID())
		s.log.Error(err)
		return nil, err
	}

	var response *counter.DecrementResponse
	proposal := newDecrementProposal(ProposalID(s.Index()), session, request, response)

	s.rsm.Proposals().Decrement().register(proposal)
	session.Proposals().Decrement().register(proposal)

	defer func() {
		session.Proposals().Decrement().unregister(proposal.ID())
		s.rsm.Proposals().Decrement().unregister(proposal.ID())
	}()

	err = s.rsm.Decrement(proposal)
	if err != nil {
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
