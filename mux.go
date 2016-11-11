package discorddotgo

import "bitbucket.org/tscs37/discorddotgo/errs"

type EventMux interface {
	Dispatch(event Event) error
	AddHandler(f interface{}) error
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

func (s *SimpleMux) AddHandler(f interface{}) error {
	log.Debug("Adding new Handler")
	switch t := f.(type) {
	case MessageHandler:
		log.Debug("Handler is full MessageHandler")
		s.AddMessageHandler(t)
	case func(Context, *Channel, *Message) error:
		log.Debug("Handler is simple MessageHandler")
		s.AddSimpleMessageHandler(SimpleMessageHandler(t))
	default:
		log.Warnf("Handler is not a supported type: %s", getTypeFrom(t))
		return errs.ErrHandlerNotSupported
	}
	return nil
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
	log.Debugf("Received Event Dispatch: %s", ev.Name())
	switch ev.Name() {
	case eventNewMessage:
		ev := ev.(*NewMessageEvent)
		ch, err := ev.Channel()
		if err != nil {
			return err
		}
		log.Debugf("Beginning execute")
		for k, v := range s.msgH {
			log.Debugf("Executing %s", v.Name())
			err := v.HandleNewMessage(ev.Context(), ch, ev.Message())
			if err != nil {
				return errs.NewHandlerError(err, k)
			}
		}
		for k, v := range s.msgSH {
			log.Debugf("Executing SH %d/%d", 1+k, len(s.msgSH))
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
