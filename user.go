package discorddgo

import "github.com/bwmarrin/discordgo"

type User struct {
	context      Context
	internalUser *discordgo.User
}

func (u *User) Username() string {
	return u.internalUser.Username
}

func (u *User) ID() string {
	return u.internalUser.ID
}

func (u *User) IsSelf() bool {
	return u.internalUser.ID == u.context.Self().ID()
}

func (u *User) IsBot() bool {
	return u.internalUser.Bot
}

func (u *User) Avatar() string {
	return u.internalUser.Avatar
}

func (u *User) Member(g *Guild) (*Member, error) {
	m, err := u.context.intSession.GuildMember(g.ID(), u.ID())
	if err != nil {
		return nil, err
	}
	return u.context.memberFromRaw(m), nil
}

type Member struct {
	context        Context
	internalMember *discordgo.Member
}

// Nick returns the Nickname of a member
func (m *Member) Nick() string {
	if m.internalMember.Nick == "" {
		return m.internalMember.User.Username
	}
	return m.internalMember.Nick
}

// SetNick sets the nickname of a user. If the user is the currently logged in
// user it will default to using @me/nick instead of the user id.
func (m *Member) SetNick(nick string) error {
	id := m.User().ID()
	if m.User().IsSelf() {
		id = "@me/nick"
	}
	guild, err := m.Guild()
	if err != nil {
		return err
	}
	return m.context.intSession.GuildMemberNickname(guild.ID(), id, nick)
}

// User returns the user of a Member
func (m *Member) User() *User {
	return m.context.userFromRaw(m.internalMember.User)
}

// Guild returns the Guild the member belongs to or an error
func (m *Member) Guild() (*Guild, error) {
	return m.context.GuildFromID(m.internalMember.GuildID)
}

// Ban will create a ban for a member or return an error
func (m *Member) Ban() error {
	return m.BanAndDeleteMessages(0)
}

// BanAndDeleteMessages will create a ban and also delete messages of the past
// days, specified by the days parameter, or return an error.
func (m *Member) BanAndDeleteMessages(days int) error {
	s := m.context.intSession
	g, err := m.Guild()
	if err != nil {
		return err
	}
	return s.GuildBanCreate(g.ID(), m.User().ID(), days)
}

func (m *Member) HasRole(r *Role) bool {
	for _, k := range m.internalMember.Roles {
		if k == r.ID() {
			return true
		}
	}
	return false
}
