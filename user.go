package discorddotgo

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

func (m *Member) Nick() string {
	return m.internalMember.Nick
}

func (m *Member) User() *User {
	return m.context.userFromRaw(m.internalMember.User)
}
