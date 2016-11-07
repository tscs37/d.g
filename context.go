package discorddotgo

import "github.com/bwmarrin/discordgo"

type Context struct {
	intSession *discordgo.Session
	exit       chan bool
}

func (c Context) int() *discordgo.Session {
	return c.intSession
}

func (c Context) ChannelFromID(id string) (*Channel, error) {
	u, err := c.intSession.Channel(id)
	if err != nil {
		return nil, err
	}
	return c.channelFromRaw(u), nil
}

func (c Context) channelFromRaw(ch *discordgo.Channel) *Channel {
	return &Channel{context: c, intChannel: ch}
}

func (c Context) UserFromID(id string) (*User, error) {
	u, err := c.intSession.User(id)
	if err != nil {
		return nil, err
	}
	return c.userFromRaw(u), nil
}

func (c Context) userFromRaw(us *discordgo.User) *User {
	return &User{context: c, internalUser: us}
}

func (c Context) GuildFromID(id string) (*Guild, error) {
	u, err := c.intSession.Guild(id)
	if err != nil {
		return nil, err
	}
	return c.guildFromRaw(u), nil
}

func (c Context) guildFromRaw(g *discordgo.Guild) *Guild {
	return &Guild{context: c, intGuild: g}
}

func (c Context) InviteFromID(id string) (*Invite, error) {
	i, err := c.int().Invite(id)
	if err != nil {
		return nil, err
	}
	return c.inviteFromRaw(i), nil
}

func (c Context) inviteFromRaw(i *discordgo.Invite) *Invite {
	return &Invite{context: c, intInvite: i}
}

func (c Context) messageFromRaw(m *discordgo.Message) *Message {
	return &Message{context: c, intMessage: m}
}

func (c Context) memberFromRaw(m *discordgo.Member) *Member {
	return &Member{context: c, internalMember: m}
}

func (c Context) roleFromRaw(r *discordgo.Role) *Role {
	return &Role{context: c, intRole: r}
}

func (c Context) Self() *User {
	u, err := c.UserFromID("@me")
	if err != nil {
		panic(err)
	}
	return u
}

func (c Context) RequestExit() {
	c.exit <- true
}
