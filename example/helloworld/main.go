package main

import (
	"fmt"
	"io/ioutil"

	ddg "go.rls.moe/d.g"
)

func main() {
	ddg.EnableDebug()

	token, err := ioutil.ReadFile("./token.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	bot, err := ddg.NewBot(string(token))
	if err != nil {
		fmt.Println(err)
		return
	}

	err = bot.SetAvatarFromURI("http://www.azquotes.com/public/pictures/authors/f9/56/f956cce770bc5d61b1d163c2e0b33a89/56163936d2b0d_rob_pike.jpg")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Ready!")

	// For Hello World we can use the SMH event sink,
	// if you want to receive other event, you'll have to write
	// your own sink.
	bot.SetSink(ddg.NewSimpleMessageHandler(handler))

	bot.BlockForExit()
}

func handler(c ddg.Context, ch *ddg.Channel, m *ddg.Message) error {
	if m.FromMe() {
		return nil
	}

	if m.DisplayText() == "ping" {
		fmt.Println("Pong!")
		m.Respond("pong")
	}

	if m.DisplayText() == "pong" {
		fmt.Println("Ping!")
		m.Respond("ping")
	}

	if m.DisplayText() == "exit" {
		c.RequestExit()
	}

	if m.DisplayText() == "nick" {

	}

	return nil
}
