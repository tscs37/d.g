package discorddotgo

import "github.com/bwmarrin/discordgo"

// Message wraps the message type with context
type Message struct {
	context    Context
	intMessage *discordgo.Message
}

func (m *Message) Text() string {
	return m.intMessage.Content
}

func (m *Message) ID() string {
	return m.intMessage.ID
}

func (m *Message) Respond(text string) (*Message, error) {
	msg, err := m.context.intSession.ChannelMessageSend(
		m.intMessage.ChannelID, text)
	return m.context.messageFromRaw(msg), err
}

func (m *Message) DisplayText() string {
	return m.intMessage.ContentWithMentionsReplaced()
}

func (m *Message) Author() *User {
	return m.context.userFromRaw(m.intMessage.Author)
}

func (m *Message) Channel() (*Channel, error) {
	ch, err := m.context.ChannelFromID(m.intMessage.ChannelID)
	if err != nil {
		return nil, err
	}
	return ch, nil
}

func (m *Message) Timestamp() string {
	return m.intMessage.Timestamp
}

func (m *Message) EditedTimestamp() string {
	return m.intMessage.EditedTimestamp
}

func (m *Message) Mentioned(u *User) bool {
	if m.intMessage.MentionEveryone {
		return true
	}

	for _, v := range m.intMessage.Mentions {
		if v.ID == u.ID() {
			return true
		}
	}

	return false
}

func (m *Message) Edit(newText string) (*Message, error) {
	ch, err := m.Channel()
	if err != nil {
		return nil, err
	}
	msg, err := m.context.int().ChannelMessageEdit(ch.ID(), m.ID(), newText)
	if err != nil {
		return nil, err
	}
	return m.context.messageFromRaw(msg), nil
}

func (m *Message) FromMe() bool {
	if m.Author().ID() == m.context.Self().ID() {
		return true
	}
	return false
}
