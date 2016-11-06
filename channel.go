package discorddotgo

import (
	"github.com/bwmarrin/discordgo"
)

// Channel wraps a channel with context
type Channel struct {
	context    *Context
	intChannel *discordgo.Channel
}

func (c *Channel) Guild() (*Guild, error) {
	g, err := c.context.intSession.Guild(c.intChannel.GuildID)
	if err != nil {
		return nil, err
	}
	return &Guild{
		context:  c.context,
		intGuild: g,
	}, nil
}

func (c *Channel) Write(message string) (*Message, error) {
	msg, err := c.context.intSession.ChannelMessageSend(c.intChannel.ID, message)
	if err != nil {
		return nil, err
	}
	return &Message{
		intMessage: msg,
		context:    c.context,
	}, nil
}
