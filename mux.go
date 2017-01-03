package discorddgo

// EventSink is a target for sending events
type EventSink interface {
	Dispatch(event Event) error
}

// SimpleMessageHandler is a very simple event
// dispatcher meant for simple hello-world type bots.
//
// It is recommended to write your own EventHandler
// that fits your need. This is much safer in terms
// of typesafety and simplicity.
type SimpleMessageHandler struct {
	msgSH func(Context, *Channel, *Message) error
}

// NewSimpleMessageHandler accepts a simple message handler
// and returns an EventSink
func NewSimpleMessageHandler(f func(Context, *Channel, *Message) error) EventSink {
	return SimpleMessageHandler{
		msgSH: f,
	}
}

func (s SimpleMessageHandler) Dispatch(ev Event) error {
	log.Debugf("Received Event Dispatch: %s", ev.Name())
	switch ev.Name() {
	case eventNewMessage:
		ev := ev.(*EventNewMessage)
		ch, err := ev.Channel()
		if err != nil {
			return err
		}
		log.Debugf("Beginning execute")
		return s.msgSH(ev.Context(), ch, ev.Message())
	default:
		return nil
	}
}
