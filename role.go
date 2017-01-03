package discorddgo

import "github.com/bwmarrin/discordgo"

type Role struct {
	context Context
	intRole *discordgo.Role
}

// Color returns the color of a role split up into the three bytes
// that make it up.
func (r *Role) Color() (byte, byte, byte) {
	red := byte(r.intRole.Color >> 16)
	green := byte(r.intRole.Color >> 8)
	blue := byte(r.intRole.Color >> 0)
	return red, green, blue
}

func (r *Role) Name() string {
	return r.intRole.Name
}

func (r *Role) ID() string {
	return r.intRole.ID
}

func (r *Role) Managed() bool {
	return r.intRole.Managed
}

func (r *Role) Hoist() bool {
	return r.intRole.Hoist
}

func (r *Role) Position() int {
	return r.intRole.Position
}
