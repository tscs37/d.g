package discorddotgo

// Handler is a generic interface for all handlers this package utilizes
type Handler interface {
	Name() string
}

// MessageHandler is an interface for any handler that accepts messages,
// outgoing or incoming
type MessageHandler interface {
	Handler
	HandleNewMessage(
		context Context,
		channel *Channel,
		message *Message) error
}

type SimpleMessageHandler func(Context, *Channel, *Message) error
