package discorddgo

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

// Message wraps the message type with context
type Message struct {
	context    Context
	intMessage *discordgo.Message
}

// Text returns the text of a message with mentions unreplaced.
func (m *Message) Text() string {
	return m.intMessage.Content
}

// ID returns the Message ID
func (m *Message) ID() string {
	return m.intMessage.ID
}

// Respond writes a message in the same channel the message originated from.
func (m *Message) Respond(text string) (*Message, error) {
	msg, err := m.context.intSession.ChannelMessageSend(
		m.intMessage.ChannelID, text)
	return m.context.messageFromRaw(msg), err
}

// DisplayText returns the text of a message with mentions replaced.
func (m *Message) DisplayText() string {
	return m.intMessage.ContentWithMentionsReplaced()
}

// AuthorUser returns the user that sent the message
func (m *Message) AuthorUser() *User {
	return m.context.userFromRaw(m.intMessage.Author)
}

// Author returns the Member that sent a message
func (m *Message) Author() (*Member, error) {
	u := m.AuthorUser()
	g, err := m.Guild()
	if err != nil {
		return nil, err
	}
	return u.Member(g)
}

// Channel returns the channel in which the message was sent or an error.
func (m *Message) Channel() (*Channel, error) {
	ch, err := m.context.ChannelFromID(m.intMessage.ChannelID)
	if err != nil {
		return nil, err
	}
	return ch, nil
}

// Guild returns the guild the message was send in. It returns an error
// if the lookup fails, for example if it's a private message channel.
func (m *Message) Guild() (*Guild, error) {
	ch, err := m.Channel()
	if err != nil {
		return nil, err
	}
	return ch.Guild()
}

// Timestamp returns the timestamp of a message
func (m *Message) Timestamp() (time.Time, error) {
	return m.intMessage.Timestamp.Parse()
}

// EditedTimestamp returns the timestamp of a message when it was edited
func (m *Message) EditedTimestamp() (time.Time, error) {
	return m.intMessage.EditedTimestamp.Parse()
}

// Mentioned returns true if the given user was mentioned.
// Note that it will not check against roles. Use MentionedMember for this.
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

// MentionedMember returns true if the specified member was mentioned.
// It might return an error in which case it will also return false.
func (m *Message) MentionedMember(member *Member) (bool, error) {
	res := m.Mentioned(member.User())
	if res {
		return true, nil
	}
	g, err := m.Guild()
	if err != nil {
		return false, err
	}
	for _, k := range m.intMessage.MentionRoles {
		role, err := m.context.RoleFromID(g.ID(), k)
		if err != nil {
			return false, err
		}
		if member.HasRole(role) {
			return true, nil
		}
	}
	return false, nil
}

// MentionedMe returns true if the current context user was mentioned.
// It might return an error in which case it will also return false.
func (m *Message) MentionedMe() (bool, error) {
	g, err := m.Guild()
	if err != nil {
		return false, err
	}
	mem, err := m.context.Self().Member(g)
	if err != nil {
		return false, err
	}
	return m.MentionedMember(mem)
}

// Edit edits the message and returns the new, edited message and an error
// if one occured.
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

// FromMe returns true if the message is from the current user.
func (m *Message) FromMe() bool {
	if m.AuthorUser().ID() == m.context.Self().ID() {
		return true
	}
	return false
}
