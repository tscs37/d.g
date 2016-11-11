package discorddotgo

import "bitbucket.org/tscs37/discorddotgo/errs"

type EventMux interface {
	Dispatch(event Event) error
}

type NewMessageMux interface {
	AddMessageHandler(m MessageHandler) error
	RemoveMessageHandler(m MessageHandler)
	AddSimpleMessageHandler(f SimpleMessageHandler)
	ResetMessageHandlers()
}

type SimpleMux struct {
	msgH  map[string]MessageHandler
	msgSH []SimpleMessageHandler
}

func (s *SimpleMux) AddMessageHandler(m MessageHandler) error {
	if s.msgH == nil {
		s.msgH = map[string]MessageHandler{}
	}
	if _, ok := s.msgH[m.Name()]; ok {
		return errs.ErrHandlerNameDuplicate
	}
	s.msgH[m.Name()] = m
	return nil
}

func (s *SimpleMux) RemoveMessageHandler(m MessageHandler) {
	delete(s.msgH, m.Name())
}

func (s *SimpleMux) AddSimpleMessageHandler(f SimpleMessageHandler) {
	s.msgSH = append(s.msgSH, f)
}

func (s *SimpleMux) ResetMessageHandlers() {
	s.msgH = map[string]MessageHandler{}
	s.msgSH = []SimpleMessageHandler{}
}

func (s *SimpleMux) Dispatch(ev Event) error {
	switch ev.Name() {
	case "new-message":
		ev := ev.(*NewMessageEvent)
		ch, err := ev.Channel()
		if err != nil {
			return err
		}
		for k, v := range s.msgH {
			err := v.HandleNewMessage(ev.Context(), ch, ev.Message())
			if err != nil {
				return errs.NewHandlerError(err, k)
			}
		}
		for _, v := range s.msgSH {
			err := v(ev.Context(), ch, ev.Message())
			if err != nil {
				return errs.NewHandlerError(err, "simple-handler")
			}
		}
		return nil
	default:
		return nil
	}
}
