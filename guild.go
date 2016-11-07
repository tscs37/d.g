package discorddotgo

import (
	"github.com/bwmarrin/discordgo"
)

type Guild struct {
	context  Context
	intGuild *discordgo.Guild
}

func (g *Guild) ID() string {
	return g.intGuild.ID
}

func (g *Guild) Channels() []*Channel {
	var ret = make([]*Channel, len(g.intGuild.Channels))
	for k, v := range g.intGuild.Channels {
		ret[k] = g.context.channelFromRaw(v)
	}
	return ret
}

func (g *Guild) Roles() []*Role {
	var ret = make([]*Role, len(g.intGuild.Roles))
	for k, v := range g.intGuild.Roles {
		ret[k] = g.context.roleFromRaw(v)
	}
	return ret
}

func (g *Guild) Members() []*Member {
	var ret = make([]*Member, len(g.intGuild.Members))
	for k, v := range g.intGuild.Members {
		ret[k] = g.context.memberFromRaw(v)
	}
	return ret
}

func (g *Guild) Icon() string {
	return g.intGuild.Icon
}
