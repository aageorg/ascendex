package main

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

func (bbo BestOrderBook) String() string {
	return fmt.Sprintf("bid price %.2f amout %f, offer price %.2f amount %f", bbo.Bid.Price, bbo.Bid.Amount, bbo.Ask.Price, bbo.Ask.Amount)
}

func ParseBBO(bbo *BestOrderBook, v interface{}) (err error) {
	bbo.Ask.Amount, err = strconv.ParseFloat(v.(map[string]any)["ask"].([]any)[1].(string), 64)
	if err != nil {
		return
	}
	bbo.Ask.Price, err = strconv.ParseFloat(v.(map[string]any)["ask"].([]any)[0].(string), 64)
	if err != nil {
		return
	}
	bbo.Bid.Amount, err = strconv.ParseFloat(v.(map[string]any)["bid"].([]any)[1].(string), 64)
	if err != nil {
		return
	}
	bbo.Bid.Price, err = strconv.ParseFloat(v.(map[string]any)["bid"].([]any)[0].(string), 64)
	if err != nil {
		return
	}
	return

}

type Message struct {
	M      string      `json:"m"`
	Symbol string      `json:"symbol,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

type Ascendex struct {
	mu     sync.Mutex
	ws_url string
	conn   *websocket.Conn
}

func NewAscendex(account_group string) *Ascendex {
	return &Ascendex{
		ws_url: "wss://ascendex.com/" + account_group + "/api/pro/v1/stream",
	}
}

func (a *Ascendex) Connection() error {
	dialer := websocket.Dialer{
		Subprotocols: []string{"json"},
	}
	c, _, err := dialer.Dial(a.ws_url, nil)
	if err != nil {
		return fmt.Errorf("Connection establishing failed: %#v\n", err)
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	a.conn = c
	return nil
}

func (a *Ascendex) Disconnect() {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.conn == nil {
		return
	}
	err := a.conn.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	a.conn = nil
}

func (a *Ascendex) SubscribeToChannel(symbol string) error {
	if a.conn == nil {
		return errors.New("Websocket connection closed")
	}
	re := regexp.MustCompile(`^[A-Z]+\_[A-Z]+$`)
	if !re.Match([]byte(symbol)) {
		return errors.New("Invalid symbol parameter")
	} else {
		symbol = strings.Replace(symbol, "_", "/", -1)
	}
	err := a.conn.WriteJSON(map[string]string{"op": "sub", "ch": "bbo:" + symbol})
	if err != nil {
		a.Disconnect()
		return err
	}
	return nil
}

// Required for testing

func (a *Ascendex) UnsubscribeFromChannel(symbol string) error {
	if a.conn == nil {
		return errors.New("Websocket connection closed")
	}
	re := regexp.MustCompile(`^[A-Z]+\_[A-Z]+$`)
	if !re.Match([]byte(symbol)) && symbol != "" {
		return errors.New("Invalid symbol parameter")
	} else if symbol != "" {
		symbol = ":" + strings.Replace(symbol, "_", "/", -1)
	}
	err := a.conn.WriteJSON(map[string]string{"op": "unsub", "ch": "bbo" + symbol})
	if err != nil {
		a.Disconnect()
		return err
	}
	return nil
}

func (a *Ascendex) ReadMessagesFromChannel(ch chan<- BestOrderBook) {
	for {
		var m Message
		err := a.conn.ReadJSON(&m)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Cannot read message from channel: "+err.Error())
		}
		var bbo BestOrderBook
		if m.M == "bbo" {
			if err = ParseBBO(&bbo, m.Data); err == nil {
				ch <- bbo
			}
		}
	}
}
func (a *Ascendex) WriteMessagesToChannel() {
	if a.conn != nil {
		a.conn.WriteJSON(map[string]string{"op": "ping"})
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: "+os.Args[0]+" pair")
		return
	}
	symb := os.Args[1]
	var asc APIClient
	asc = NewAscendex("0")

	err := asc.Connection()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	go func() {
		for {
			asc.WriteMessagesToChannel()
			time.Sleep(15 * time.Second)
		}
	}()
	err = asc.SubscribeToChannel(symb)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	ch := make(chan BestOrderBook, 30)
	go asc.ReadMessagesFromChannel(ch)
	for bbo := range ch {
		fmt.Fprintln(os.Stdout, bbo)
	}
}
