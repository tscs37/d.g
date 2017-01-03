package discorddgo

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"strings"

	"os"

	"github.com/bwmarrin/discordgo"
)

// A bot is a running discord session intended for usage
// with a Bot-type Login
type Bot struct {
	context Context
	evMux   EventSink
}

// NewBot uses the given token to log into discord and return
// a Bot-type session.
//
// The "Bot " authorization string is automatically appended
// if not found.
func NewBot(token string) (*Bot, error) {
	if !strings.HasPrefix(token, "Bot ") {
		token = "Bot " + token
	}
	dg, err := discordgo.New()
	if err != nil {
		return nil, err
	}
	b := &Bot{
		context: Context{intSession: dg, exit: make(chan bool)},
		evMux:   nil,
	}
	dg.AddHandler(b.execHandler)

	err = dg.Open()
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Sink returns the current EventSink
//
// This may be used by plugins to dispatch custom events.
func (b *Bot) Sink() EventSink {
	return b.evMux
}

// SetSink sets the EventSink to be used by D.G
func (b *Bot) SetSink(ev EventSink) {
	b.evMux = ev
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
		return b.evMux.Dispatch(&EventNewMessage{context: b.context, m: t})
	case *discordgo.MessageUpdate:
		return b.evMux.Dispatch(&EventMessageUpdate{context: b.context, m: t})
	case *discordgo.TypingStart:
		return b.evMux.Dispatch(&EventUserTyping{context: b.context, ev: t})
	}
	return nil
}

// SetAvatarFromReader will read the given io.Reader and
// set the content as avatar
//
// Remember to close the reader after use if applicable.
func (b *Bot) SetAvatarFromReader(image io.Reader) error {
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

// SetAvatarFromURI will download the URI data and set it as avatar
func (b *Bot) SetAvatarFromURI(uri string) error {
	resp, err := http.Get(uri)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return b.SetAvatarFromReader(resp.Body)
}

// SetAvatarFromFile will read the file and set it as avatar
func (b *Bot) SetAvatarFromFile(file string) error {
	dat, err := os.Open(file)
	if err != nil {
		return err
	}

	defer dat.Close()

	return b.SetAvatarFromReader(dat)
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
