// Code generated by atomix-go-framework. DO NOT EDIT.
package list

import (
	list "github.com/atomix/atomix-api/go/atomix/primitive/list"
	"github.com/atomix/atomix-go-framework/pkg/atomix/errors"
	"github.com/atomix/atomix-go-framework/pkg/atomix/logging"
	"github.com/atomix/atomix-go-framework/pkg/atomix/storage/protocol/rsm"
	"github.com/atomix/atomix-go-framework/pkg/atomix/util"
	"github.com/golang/protobuf/proto"
	"io"
)

const Type = "List"

const (
	sizeOp     = "Size"
	appendOp   = "Append"
	insertOp   = "Insert"
	getOp      = "Get"
	setOp      = "Set"
	removeOp   = "Remove"
	clearOp    = "Clear"
	eventsOp   = "Events"
	elementsOp = "Elements"
)

var newServiceFunc rsm.NewServiceFunc

func registerServiceFunc(rsmf NewServiceFunc) {
	newServiceFunc = func(scheduler rsm.Scheduler, context rsm.ServiceContext) rsm.Service {
		service := &ServiceAdaptor{
			Service: rsm.NewService(scheduler, context),
			rsm:     rsmf(newServiceContext(scheduler)),
			log:     logging.GetLogger("atomix", "list", "service"),
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
	s.RegisterUnaryOperation(sizeOp, s.size)
	s.RegisterUnaryOperation(appendOp, s.append)
	s.RegisterUnaryOperation(insertOp, s.insert)
	s.RegisterUnaryOperation(getOp, s.get)
	s.RegisterUnaryOperation(setOp, s.set)
	s.RegisterUnaryOperation(removeOp, s.remove)
	s.RegisterUnaryOperation(clearOp, s.clear)
	s.RegisterStreamOperation(eventsOp, s.events)
	s.RegisterStreamOperation(elementsOp, s.elements)
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
	state := &ListState{}
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
func (s *ServiceAdaptor) size(input []byte, rsmSession rsm.Session) ([]byte, error) {
	request := &list.SizeRequest{}
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

	var response *list.SizeResponse
	proposal := newSizeProposal(ProposalID(s.Index()), session, request, response)

	s.rsm.Proposals().Size().register(proposal)
	session.Proposals().Size().register(proposal)

	defer func() {
		session.Proposals().Size().unregister(proposal.ID())
		s.rsm.Proposals().Size().unregister(proposal.ID())
	}()

	err = s.rsm.Size(proposal)
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
func (s *ServiceAdaptor) append(input []byte, rsmSession rsm.Session) ([]byte, error) {
	request := &list.AppendRequest{}
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

	var response *list.AppendResponse
	proposal := newAppendProposal(ProposalID(s.Index()), session, request, response)

	s.rsm.Proposals().Append().register(proposal)
	session.Proposals().Append().register(proposal)

	defer func() {
		session.Proposals().Append().unregister(proposal.ID())
		s.rsm.Proposals().Append().unregister(proposal.ID())
	}()

	err = s.rsm.Append(proposal)
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
func (s *ServiceAdaptor) insert(input []byte, rsmSession rsm.Session) ([]byte, error) {
	request := &list.InsertRequest{}
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

	var response *list.InsertResponse
	proposal := newInsertProposal(ProposalID(s.Index()), session, request, response)

	s.rsm.Proposals().Insert().register(proposal)
	session.Proposals().Insert().register(proposal)

	defer func() {
		session.Proposals().Insert().unregister(proposal.ID())
		s.rsm.Proposals().Insert().unregister(proposal.ID())
	}()

	err = s.rsm.Insert(proposal)
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
	request := &list.GetRequest{}
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

	var response *list.GetResponse
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
func (s *ServiceAdaptor) set(input []byte, rsmSession rsm.Session) ([]byte, error) {
	request := &list.SetRequest{}
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

	var response *list.SetResponse
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
func (s *ServiceAdaptor) remove(input []byte, rsmSession rsm.Session) ([]byte, error) {
	request := &list.RemoveRequest{}
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

	var response *list.RemoveResponse
	proposal := newRemoveProposal(ProposalID(s.Index()), session, request, response)

	s.rsm.Proposals().Remove().register(proposal)
	session.Proposals().Remove().register(proposal)

	defer func() {
		session.Proposals().Remove().unregister(proposal.ID())
		s.rsm.Proposals().Remove().unregister(proposal.ID())
	}()

	err = s.rsm.Remove(proposal)
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
func (s *ServiceAdaptor) clear(input []byte, rsmSession rsm.Session) ([]byte, error) {
	request := &list.ClearRequest{}
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

	var response *list.ClearResponse
	proposal := newClearProposal(ProposalID(s.Index()), session, request, response)

	s.rsm.Proposals().Clear().register(proposal)
	session.Proposals().Clear().register(proposal)

	defer func() {
		session.Proposals().Clear().unregister(proposal.ID())
		s.rsm.Proposals().Clear().unregister(proposal.ID())
	}()

	err = s.rsm.Clear(proposal)
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
func (s *ServiceAdaptor) events(input []byte, rsmSession rsm.Session, stream rsm.Stream) (rsm.StreamCloser, error) {
	request := &list.EventsRequest{}
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

	proposal := newEventsProposal(ProposalID(stream.ID()), session, request, stream)

	s.rsm.Proposals().Events().register(proposal)
	session.Proposals().Events().register(proposal)

	err = s.rsm.Events(proposal)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}
	return func() {
		session.Proposals().Events().unregister(proposal.ID())
		s.rsm.Proposals().Events().unregister(proposal.ID())
	}, nil
}

func (s *ServiceAdaptor) elements(input []byte, rsmSession rsm.Session, stream rsm.Stream) (rsm.StreamCloser, error) {
	request := &list.ElementsRequest{}
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

	proposal := newElementsProposal(ProposalID(stream.ID()), session, request, stream)

	s.rsm.Proposals().Elements().register(proposal)
	session.Proposals().Elements().register(proposal)

	err = s.rsm.Elements(proposal)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}
	return func() {
		session.Proposals().Elements().unregister(proposal.ID())
		s.rsm.Proposals().Elements().unregister(proposal.ID())
	}, nil
}

var _ rsm.Service = &ServiceAdaptor{}
