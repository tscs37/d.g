package discorddotgo

import (
	"github.com/bwmarrin/discordgo"
)

type Guild struct {
	context  *Context
	intGuild *discordgo.Guild
}

func (g *Guild) ID() string {
	return g.intGuild.ID
}
