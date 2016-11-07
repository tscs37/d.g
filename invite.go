package discorddotgo

import "github.com/bwmarrin/discordgo"

type Invite struct {
	context   Context
	intInvite *discordgo.Invite
}

func (i *Invite) Accept() error {
	_, err := i.context.int().InviteAccept(i.intInvite.Code)
	return err
}

func (i *Invite) Guild() *Guild {
	return i.context.guildFromRaw(i.intInvite.Guild)
}

func (i *Invite) Channel() *Channel {
	return i.context.channelFromRaw(i.intInvite.Channel)
}
