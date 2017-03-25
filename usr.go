package discorddgo

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"os"

	"github.com/bwmarrin/discordgo"
)

// A bot is a running discord session intended for usage
// with a Bot-type Login
type UserSession struct {
	context Context
	evMux   EventSink
}

// NewUser uses the given token to log into discord and return
// a User-type session.
func NewUser(token string) (*UserSession, error) {
	dg, err := discordgo.New(token)
	if err != nil {
		return nil, err
	}
	b := &UserSession{
		context: Context{intSession: dg, exit: nil},
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
func (b *UserSession) Sink() EventSink {
	return b.evMux
}

// SetSink sets the EventSink to be used by D.G
func (b *UserSession) SetSink(ev EventSink) {
	b.evMux = ev
}

// CurrentContext returns the current context pointer.
func (b *UserSession) CurrentContext() Context {
	return b.context
}

// Me returns the current bot user (using @me) or an error.
func (b *UserSession) Me() (*User, error) {
	return b.context.UserFromID("@me")
}

func (b *UserSession) execHandler(s *discordgo.Session, ev interface{}) {
	switch t := ev.(type) {
	case *discordgo.MessageCreate:
		b.evMux.Dispatch(&EventNewMessage{context: b.context, m: t})
	case *discordgo.MessageUpdate:
		b.evMux.Dispatch(&EventMessageUpdate{context: b.context, m: t})
	case *discordgo.TypingStart:
		b.evMux.Dispatch(&EventUserTyping{context: b.context, ev: t})
	}
}

// SetAvatarFromReader will read the given io.Reader and
// set the content as avatar
//
// Remember to close the reader after use if applicable.
func (b *UserSession) SetAvatarFromReader(image io.Reader) error {
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
func (b *UserSession) SetAvatarFromURI(uri string) error {
	resp, err := http.Get(uri)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return b.SetAvatarFromReader(resp.Body)
}

// SetAvatarFromFile will read the file and set it as avatar
func (b *UserSession) SetAvatarFromFile(file string) error {
	dat, err := os.Open(file)
	if err != nil {
		return err
	}

	defer dat.Close()

	return b.SetAvatarFromReader(dat)
}

// GetDiscordgoSession returns the underlying session the bot instance is using.
// This is not recommended unless the functionality you want to use is not
// implemented on d.g yet.
func (b *UserSession) GetDiscordgoSession() *discordgo.Session {
	return b.context.intSession
}
