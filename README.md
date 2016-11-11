# discord.go

discord.go is a wrapper for bwmarrin/discordgo to enable a few "oh that's neat"
functions.

d.g is heavily context-based and focues on bot development.

The goal is that all functionality a bot might need is provided out of the box
and can be used without having to think about the discord API.

**Note** The official import path is `pit.timschuster.info/d.g`, this URL
will point at the newest d.g on a best effort basis.

## How to create a bot

A bot is easily created with d.g

Here is an example:

```go
package main

import ddg "bitbucket.org/tscs37/discorddotgo"

func main() {
    bot, err := ddg.NewBot("<your token>")
    if err != nil {
        println(err)
        return
    }

    bot.AddSimpleMessageHandler(handler)

    bot.BlockForExit()
}

func handler(c ddg.Context, ch *ddg.Channel, m *ddg.Message) error {
    if m.FromMe() {
        return nil
    }

    if m.DisplayText() == "ping" {
        m.Respond("pong")
    }

    if m.DisplayText() == "exit" {
        c.RequestExit()
    }

    return nil
}
```

Refer to the godoc documentation for further information.

(The Documentation is still being worked on, the library is still under development)

## Project Goals

This library aims to completely wrap discordgo.

The first step is to wrap all read-only functionality and some
basic response methods to allow bots interacting with the chat.

The second step is to implement all management methods so that easy
functions exist to assign roles, ban users or create channels.

The application using this library should not have to come in contact
with any discordgo types but the library should also provide access to
the underlying session so they *can* use the original discordgo package if
needed.