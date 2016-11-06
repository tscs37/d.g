package discorddotgo

import (
	"bitbucket.org/tscs37/discorddotgo/errs"
	"github.com/bwmarrin/discordgo"
)



type Bot struct {
	context *Context
	msgH    map[string]MessageHandler
}

func NewBot(token string) (*Bot, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	b := &Bot{
		context: &Context{intSession: dg},
		msgH:    map[string]MessageHandler{},
	}

	dg.AddHandler(b.execMessageHandlers)

	err = dg.Open()
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Add a message handler to the global context
func (b *Bot) AddMessageHandler(m MessageHandler) error {
	if _, ok := b.msgH[m.Name()]; ok {
		return errs.ErrHandlerNameDuplicate
	}
	b.msgH[m.Name()] = m
	return nil
}

// RemoveMessageHandler will remove a message handler from the global
// context
func (b *Bot) RemoveMessageHandler(m MessageHandler) {
	delete(b.msgH, m.Name())
}

// ResetMessageHandlers will delete all registered message handlers
func (b *Bot) ResetMessageHandlers() {
	b.msgH = map[string]MessageHandler{}
}

// CurrentContext returns the current context pointer.
func (b *Bot) CurrentContext() *Context {
	return b.context
}

func (b *Bot) execMessageHandlers(
	s *discordgo.Session, m *discordgo.Message) error {
	rawChan, err := s.Channel(m.ChannelID)
	if err != nil {
		return err
	}
	ch := &Channel{intChannel: rawChan, context: b.context}
	msg := &Message{intMessage: m, context: b.context}
	for k, v := range b.msgH {
		err := v.HandleNewMessage(b.context, ch, msg)
		if err != nil {
			return errs.NewHandlerError(err, k)
		}
	}
	return nil
}
