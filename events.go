package discorddotgo

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

type NewMessageEvent struct {
	context Context
	m       *discordgo.MessageCreate
}

func (n NewMessageEvent) Name() string     { return eventNewMessage }
func (n NewMessageEvent) Context() Context { return n.context }
func (n *NewMessageEvent) Message() *Message {
	return n.context.messageFromRaw(n.m.Message)
}
func (n *NewMessageEvent) Channel() (*Channel, error) {
	return n.context.ChannelFromID(n.m.ChannelID)
}

type MessageUpdateEvent struct {
	context Context
	msg     *discordgo.MessageUpdate
}

func (n MessageUpdateEvent) Name() string     { return eventMessageUpdate }
func (n MessageUpdateEvent) Context() Context { return n.context }

type UserTypingEvent struct {
	context Context
	ev      *discordgo.TypingStart
}

func (n UserTypingEvent) Name() string     { return eventUserTyping }
func (n UserTypingEvent) Context() Context { return n.context }
func (n *UserTypingEvent) User() (*User, error) {
	return n.context.UserFromID(n.ev.UserID)
}

func (n *UserTypingEvent) Channel() (*Channel, error) {
	return n.context.ChannelFromID(n.ev.ChannelID)
}
