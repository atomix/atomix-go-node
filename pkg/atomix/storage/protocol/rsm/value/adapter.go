// Code generated by atomix-go-framework. DO NOT EDIT.
package value

import (
	value "github.com/atomix/atomix-api/go/atomix/primitive/value"
	"github.com/atomix/atomix-go-framework/pkg/atomix/errors"
	"github.com/atomix/atomix-go-framework/pkg/atomix/logging"
	"github.com/atomix/atomix-go-framework/pkg/atomix/storage/protocol/rsm"
	"github.com/golang/protobuf/proto"
	"io"
)

var log = logging.GetLogger("atomix", "value", "service")

const Type = "Value"

const (
	setOp    = "Set"
	getOp    = "Get"
	eventsOp = "Events"
)

var newServiceFunc rsm.NewServiceFunc

func registerServiceFunc(rsmf NewServiceFunc) {
	newServiceFunc = func(scheduler rsm.Scheduler, context rsm.ServiceContext) rsm.Service {
		service := &ServiceAdaptor{
			Service: rsm.NewService(scheduler, context),
			rsm:     rsmf(newServiceContext(scheduler)),
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
}

func (s *ServiceAdaptor) init() {
	s.RegisterUnaryOperation(setOp, s.set)
	s.RegisterUnaryOperation(getOp, s.get)
	s.RegisterStreamOperation(eventsOp, s.events)
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
	err := s.rsm.Backup(newSnapshotWriter(writer))
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (s *ServiceAdaptor) Restore(reader io.Reader) error {
	err := s.rsm.Restore(newSnapshotReader(reader))
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
func (s *ServiceAdaptor) set(input []byte, rsmSession rsm.Session) ([]byte, error) {
	request := &value.SetRequest{}
	err := proto.Unmarshal(input, request)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	session, ok := s.rsm.Sessions().Get(SessionID(rsmSession.ID()))
	if !ok {
		err := errors.NewConflict("session %d not found", rsmSession.ID())
		log.Warn(err)
		return nil, err
	}

	proposal := newSetProposal(ProposalID(s.Index()), session, request)

	s.rsm.Proposals().Set().register(proposal)
	session.Proposals().Set().register(proposal)

	defer func() {
		session.Proposals().Set().unregister(proposal.ID())
		s.rsm.Proposals().Set().unregister(proposal.ID())
	}()

	log.Debugf("Proposing SetProposal %s", proposal)
	err = s.rsm.Set(proposal)
	if err != nil {
		log.Warn(err)
		return nil, err
	}

	output, err := proto.Marshal(proposal.response())
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return output, nil
}
func (s *ServiceAdaptor) get(input []byte, rsmSession rsm.Session) ([]byte, error) {
	request := &value.GetRequest{}
	err := proto.Unmarshal(input, request)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	session, ok := s.rsm.Sessions().Get(SessionID(rsmSession.ID()))
	if !ok {
		err := errors.NewConflict("session %d not found", rsmSession.ID())
		log.Warn(err)
		return nil, err
	}

	proposal := newGetProposal(ProposalID(s.Index()), session, request)

	s.rsm.Proposals().Get().register(proposal)
	session.Proposals().Get().register(proposal)

	defer func() {
		session.Proposals().Get().unregister(proposal.ID())
		s.rsm.Proposals().Get().unregister(proposal.ID())
	}()

	log.Debugf("Proposing GetProposal %s", proposal)
	err = s.rsm.Get(proposal)
	if err != nil {
		log.Warn(err)
		return nil, err
	}

	output, err := proto.Marshal(proposal.response())
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return output, nil
}
func (s *ServiceAdaptor) events(input []byte, rsmSession rsm.Session, stream rsm.Stream) (rsm.StreamCloser, error) {
	request := &value.EventsRequest{}
	err := proto.Unmarshal(input, request)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	session, ok := s.rsm.Sessions().Get(SessionID(rsmSession.ID()))
	if !ok {
		err := errors.NewConflict("session %d not found", rsmSession.ID())
		log.Warn(err)
		return nil, err
	}

	proposal := newEventsProposal(ProposalID(s.Index()), session, request, stream)

	s.rsm.Proposals().Events().register(proposal)
	session.Proposals().Events().register(proposal)

	log.Debugf("Proposing EventsProposal %s", proposal)
	err = s.rsm.Events(proposal)
	if err != nil {
		log.Warn(err)
		return nil, err
	}
	return func() {
		session.Proposals().Events().unregister(proposal.ID())
		s.rsm.Proposals().Events().unregister(proposal.ID())
	}, nil
}

var _ rsm.Service = &ServiceAdaptor{}
