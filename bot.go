package discorddotgo

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	context Context
	evMux   EventMux
}

func NewBot(token string) (*Bot, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	b := &Bot{
		context: Context{intSession: dg, exit: make(chan bool)},
		evMux:   &SimpleMux{},
	}
	b.evMux.(*SimpleMux).ResetMessageHandlers()
	dg.AddHandler(b.execHandler)

	err = dg.Open()
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Mux returns the currently used Event Mux of the bot
func (b *Bot) Mux() EventMux {
	return b.evMux
}

// SetMux sets a custom event mux/dispatch
func (b *Bot) SetMux(ev EventMux) {
	b.evMux = ev
}

// RestoreDefaultMux sets the default SimpleMux handler
func (b *Bot) RestoreDefaultMux() {
	b.evMux = &SimpleMux{}
	b.evMux.(*SimpleMux).ResetMessageHandlers()
}

// CurrentContext returns the current context pointer.
func (b *Bot) CurrentContext() Context {
	return b.context
}

// Me returns the current bot user (using @me) or an error.
func (b *Bot) Me() (*User, error) {
	return b.context.UserFromID("@me")
}

func (b *Bot) execHandler(s *discordgo.Session, ev interface{}) error {
	switch t := ev.(type) {
	case *discordgo.MessageCreate:
		return b.evMux.Dispatch(&NewMessageEvent{context: b.context, m: t})
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

// BlockForExit will block on the exit channel of the bot.
// Any code that has access to the application context can unblock
// this method, which usually means exiting the application.
// Exiting the entire application is not enforced, this function only blocks,
// so it may be used to restart the bot too.
func (b *Bot) BlockForExit() {
	<-b.context.exit
}

// GetDiscordgoSession returns the underlying session the bot instance is using.
// This is not recommended unless the functionality you want to use is not
// implemented on d.g yet.
func (b *Bot) GetDiscordgoSession() *discordgo.Session {
	return b.context.intSession
}
