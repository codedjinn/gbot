package main

import (
	"fmt"
	"net"
	"gopkg.in/irc.v3"
)

type WorldTimeBody struct {
	datetime string
}

//var wg sync.WaitGroup

func main() {

	clientUri := "chat.freenode.net:6667"
	conn, err := net.Dial("tcp", clientUri)
	if (err != nil) {
		fmt.Println("Cannot connect to " + clientUri)
		return
	}
	
	config := irc.ClientConfig{
		Nick: "i_have_a_nick",
		Pass: "password",
		User: "username",
		Name: "Full Name",
		Handler: irc.HandlerFunc(func(c *irc.Client, m *irc.Message) {
			if m.Command == "001" {
				// 001 is a welcome event, so we join channels there
				c.Write("JOIN #bot-test-chan")
			} else if m.Command == "PRIVMSG" && c.FromChannel(m) {
				// Create a handler on all messages.
				c.WriteMessage(&irc.Message{
					Command: "PRIVMSG",
					Params: []string{
						m.Params[0],
						m.Trailing(),
					},
				})
			}
		}),
	}

	client := irc.NewClient(conn, config)
	err1 := client.Run()
	if err1 != nil {
		fmt.Println("FUCK")
	}

	fmt.Println("GBOT")
	fmt.Println("----")

	bot := NewBot("http://worldtimeapi.org/api")
	go bot.Initialize()

	bot.Process()

	fmt.Println()
	fmt.Println("done")
}
