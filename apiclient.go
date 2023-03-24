package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"os"
	"time"
)

type Data struct {
	Ts  int64    `json:"ts"`
	Bid []string `json:"bid"`
	Ask []string `json:"ask"`
}

type Message struct {
	M      string `json:"m"`
	Symbol string `json:"symbol"`
	Data   Data   `json:"data"`
}

type Client struct {
	ws_url    string
	conn      *websocket.Conn
	is_closed bool
}

func (cl *Client) Connection() error {
	dialer := websocket.Dialer{
		Subprotocols: []string{"json"},
	}
	c, _, err := dialer.Dial(cl.ws_url, nil)
	if err != nil {
		return fmt.Errorf("Connection establishing failed: %#v\n", err)
	}
	cl.conn = c
	cl.is_closed = false
	return nil
}

func (cl *Client) Disconnect() error {
	err := cl.conn.Close()
	if err != nil {
		return err
	}
	cl.is_closed = true
	return nil
}

func (cl *Client) SubscribeToChannel(symbol string) error {
	if symbol != "" {
		symbol = ":" + symbol
	}
	err := cl.conn.WriteJSON(map[string]any{"op": "sub", "ch": "bbo" + symbol})
	if err != nil {
		cl.Disconnect()
		return err
	}
	return nil
}

func (cl *Client) ReadMessagesFromChannel(ch chan Message) {
	var m Message
	for {
		err := cl.conn.ReadJSON(&m)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Cannot read message from channel: "+err.Error())
			cl.Disconnect()
			return
		}
		if m.M == "bbo" {
			ch <- m
		}
	}
}

func (cl *Client) WriteMessagesToChannel() {
	for {
		cl.conn.WriteJSON(map[string]any{"op": "ping"})
		time.Sleep(15 * time.Second)
	}
}
