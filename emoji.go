package discorddgo

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type Emoji struct {
	context  Context
	intEmoji discordgo.Emoji
}

func (e *Emoji) String() string {
	format := ":%s:"
	if !e.intEmoji.RequireColons {
		format = "%s"
	}
	return fmt.Sprintf(format, e.intEmoji.ID)
}

func (e *Emoji) Name() string {
	return e.intEmoji.Name
}

func (e *Emoji) Managed() bool {
	return e.intEmoji.Managed
}
