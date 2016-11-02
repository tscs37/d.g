package discorddotgo

import (
	"bitbucket.org/tscs37/discorddotgo/errs"
	"github.com/bwmarrin/discordgo"
)

type Handler interface {
	Name() string
}

type MessageHandler interface {
	Handler
	HandleNewMessage(
		context *Context,
		channel *Channel,
		message *Message) error
}

type Channel struct {
	c  *Context
	ch *discordgo.Channel
}

type Message struct {
	c   *Context
	msg *discordgo.Message
}

func (m *Message) Text() string {
	return m.msg.Content
}

func (m *Message) Respond(text string) (*Message, error) {
	msg, err := m.c.s.ChannelMessageSend(m.msg.ChannelID, text)
	return &Message{msg: msg, c: m.c}, err
}

type Context struct {
	s *discordgo.Session
}

type Bot struct {
	c    *Context
	msgH map[string]MessageHandler
}

func NewBot(token string) (*Bot, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	b := &Bot{
		c:    &Context{s: dg},
		msgH: map[string]MessageHandler{},
	}

	dg.AddHandler(b.execMessageHandlers)

	err = dg.Open()
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (b *Bot) AddMessageHandler(m MessageHandler) error {
	if _, ok := b.msgH[m.Name()]; ok {
		return errs.ErrHandlerNameDuplicate
	}
	b.msgH[m.Name()] = m
	return nil
}

func (b *Bot) execMessageHandlers(
	s *discordgo.Session, m *discordgo.Message) error {
	rawChan, err := s.Channel(m.ChannelID)
	if err != nil {
		return err
	}
	ch := &Channel{ch: rawChan, c: b.c}
	msg := &Message{msg: m, c: b.c}
	for k, v := range b.msgH {
		err := v.HandleNewMessage(b.c, ch, msg)
		if err != nil {
			return errs.NewHandlerError(err, k)
		}
	}
	return nil
}
