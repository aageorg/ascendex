package main

import (
	"fmt"
	"os"
)

func (m Message) String() string {
	return fmt.Sprintf("%s: bid %s offer %s ", m.Symbol, m.Data.Bid[0], m.Data.Ask[0])
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: "+os.Args[0]+" pair")
		return
	}
	symb := os.Args[1]
	cl := Client{
		ws_url: "wss://ascendex.com/0/api/pro/v1/stream",
	}
	err := cl.Connection()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	go cl.WriteMessagesToChannel()

	err = cl.SubscribeToChannel(symb)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	ch := make(chan Message, 30)
	go cl.ReadMessagesFromChannel(ch)
	for m := range ch {
		fmt.Fprintln(os.Stdout, m)
	}
}
