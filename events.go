package discorddgo

import "github.com/bwmarrin/discordgo"

const (
	eventNewMessage    = "new-message"
	eventMessageUpdate = "message-update"
	eventUserTyping    = "user-typing"
)

type Event interface {
	Name() string
	Context() Context
}

type EventNewMessage struct {
	context Context
	m       *discordgo.MessageCreate
}

func (n EventNewMessage) Name() string     { return eventNewMessage }
func (n EventNewMessage) Context() Context { return n.context }
func (n *EventNewMessage) Message() *Message {
	return n.context.messageFromRaw(n.m.Message)
}
func (n *EventNewMessage) Channel() (*Channel, error) {
	return n.context.ChannelFromID(n.m.ChannelID)
}

type EventMessageUpdate struct {
	context Context
	m       *discordgo.MessageUpdate
}

func (n EventMessageUpdate) Name() string     { return eventMessageUpdate }
func (n EventMessageUpdate) Context() Context { return n.context }
func (n EventMessageUpdate) Message() *Message {
	return n.context.messageFromRaw(n.m.Message)
}
func (n *EventMessageUpdate) Channel() (*Channel, error) {
	return n.context.ChannelFromID(n.m.ChannelID)
}

type EventUserTyping struct {
	context Context
	ev      *discordgo.TypingStart
}

func (n EventUserTyping) Name() string     { return eventUserTyping }
func (n EventUserTyping) Context() Context { return n.context }
func (n EventUserTyping) User() (*User, error) {
	return n.context.UserFromID(n.ev.UserID)
}

func (n EventUserTyping) Channel() (*Channel, error) {
	return n.context.ChannelFromID(n.ev.ChannelID)
}
