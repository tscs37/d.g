/*
discord.go is a wrapper for bwmarrin/discordgo to enable a few "oh that's neat"
functions.

d.g is heavily context-based and focues on bot development.

The goal is that all functionality a bot might need is provided out of the box
and can be used without having to think about the discord API.

To start a new bot, simply call NewBot with the token for your bot
and you're basically good to go.

An example bot is provided in example/helloworld
*/
package discorddgo
