package discorddotgo

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"bitbucket.org/tscs37/discorddotgo/errs"
	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	context Context
	msgH    map[string]MessageHandler
	msgSH   []SimpleMessageHandler
}

func NewBot(token string) (*Bot, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	b := &Bot{
		context: Context{intSession: dg, exit: make(chan bool)},
		msgH:    map[string]MessageHandler{},
		msgSH:   []SimpleMessageHandler{},
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

func (b *Bot) AddSimpleMessageHandler(f SimpleMessageHandler) {
	b.msgSH = append(b.msgSH, f)
}

// RemoveMessageHandler will remove a message handler from the global
// context
func (b *Bot) RemoveMessageHandler(m MessageHandler) {
	delete(b.msgH, m.Name())
}

// ResetMessageHandlers will delete all registered message handlers
func (b *Bot) ResetMessageHandlers() {
	b.msgH = map[string]MessageHandler{}
	b.msgSH = []SimpleMessageHandler{}
}

// CurrentContext returns the current context pointer.
func (b *Bot) CurrentContext() Context {
	return b.context
}

func (b *Bot) Me() (*User, error) {
	return b.context.UserFromID("@me")
}

func (b *Bot) execMessageHandlers(
	s *discordgo.Session, m *discordgo.MessageCreate) error {
	rawChan, err := s.Channel(m.ChannelID)
	if err != nil {
		return err
	}
	ch := &Channel{intChannel: rawChan, context: b.context}
	msg := &Message{intMessage: m.Message, context: b.context}
	for k, v := range b.msgH {
		err := v.HandleNewMessage(b.context, ch, msg)
		if err != nil {
			return errs.NewHandlerError(err, k)
		}
	}
	for _, v := range b.msgSH {
		err := v(b.context, ch, msg)
		if err != nil {
			return errs.NewHandlerError(err, "simple-handler")
		}
	}
	return nil
}

func (b *Bot) SetAvatarFromFile(image io.Reader) error {
	data, err := ioutil.ReadAll(image)
	if err != nil {
		return err
	}

	me, err := b.Me()
	if err != nil {
		return err
	}

	b64 := base64.StdEncoding.EncodeToString(data)

	avatar := fmt.Sprintf("data:%s;base64,%s", http.DetectContentType(data), b64)

	_, err = b.context.intSession.UserUpdate("", "", me.internalUser.Username, avatar, "")
	return err
}

func (b *Bot) SetAvatarFromURI(uri string) error {
	resp, err := http.Get(uri)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return b.SetAvatarFromFile(resp.Body)
}

func (b *Bot) BlockForExit() {
	<-b.context.exit
}
