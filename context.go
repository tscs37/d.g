package discorddgo

import "github.com/bwmarrin/discordgo"

// Context represents a particular Discord Session. It is available to all
// objects it spawns. If operating a user, it will usually contain a context
// and give all objects it creates the context too. This allows to
// access the DiscordSession from everywhere, thus allowing to use any method of
// a type in any code context.
//
// An example would be to ban a user that has a mention in a message received.
// Simply retrieve the mention, get the user object from this mention and
// call the Ban() method on the user.
type Context struct {
	intSession *discordgo.Session
	exit       chan bool
}

func (c Context) int() *discordgo.Session {
	return c.intSession
}

// ChannelFromID converts a ChannelID to a Channel instance or returns an error
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

// UserFromID converts a UserID to a User instance or returns an error
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

// GuildFromID converts a GuildID to a Guild instance or return an error
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

// InviteFromID converts a InviteID to an Invite instance or returns an error
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

func (c Context) RoleFromID(gid, id string) (*Role, error) {
	r, err := c.int().State.Role(gid, id)
	if err != nil {
		return nil, err
	}
	return &Role{intRole: r, context: c}, nil
}

// Self returns the Context user. If the context has no user, it panics.
func (c Context) Self() *User {
	u, err := c.UserFromID("@me")
	if err != nil {
		log.Fatal("I am not a user!")
	}
	return u
}

// RequestExit will send a value to the exit channel. If your bot is blocking
// on BlockForExit() it will return from this function. Beware that this function
// will block if the bot is not waiting on it.
func (c Context) RequestExit() {
	c.exit <- true
}
