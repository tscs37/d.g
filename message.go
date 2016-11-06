package discorddotgo

import "github.com/bwmarrin/discordgo"

// Message wraps the message type with context
type Message struct {
	context    *Context
	intMessage *discordgo.Message
}

func (m *Message) Text() string {
	return m.intMessage.Content
}

func (m *Message) Respond(text string) (*Message, error) {
	msg, err := m.context.intSession.ChannelMessageSend(
		m.intMessage.ChannelID, text)
	return &Message{intMessage: msg, context: m.context}, err
}
