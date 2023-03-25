package main

import (
	"context"
	"testing"
	"time"
)

func ConnectedClient() *Ascendex {
	a := NewAscendex("0")
	a.Connection()
	return a
}
func EmptyClient() *Ascendex {
	return &Ascendex{}
}

var a = ConnectedClient()
var b = EmptyClient()

func TestConnection_Success(t *testing.T) {
	asc := NewAscendex("0")
	err := asc.Connection()
	defer asc.Disconnect()
	if err != nil {
		t.FailNow()
	}
}

func TestConnection_Failure(t *testing.T) {
	asc := &Ascendex{ws_url: "invalid"}
	err := asc.Connection()
	if err == nil {
		t.FailNow()
	}
}

func TestSubscribeToChannel_Success(t *testing.T) {
	err := a.SubscribeToChannel("BTC_USDT")
	defer a.UnsubscribeFromChannel("BTC_USDT")
	if err != nil {
		t.FailNow()
	}

}

func TestSubscribeToChannel_Failure(t *testing.T) {
	err := b.SubscribeToChannel("BTC_USDT")
	if err == nil {
		t.FailNow()
	}
}

func TestWriteMessagesToChannel_Success(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	ch := make(chan bool)
	go func() {
		for {
			if a.conn == nil {
				continue
			}
			var m Message
			a.conn.ReadJSON(&m)
			if m.M == "pong" {
				ch <- true
				break
			}
		}
	}()
	a.WriteMessagesToChannel()
	select {
	case <-ch:
		return
	case <-ctx.Done():
		t.FailNow()
	}
}


func TestReadMessagesFromChannel_Success(t *testing.T) {
	a.SubscribeToChannel("BTC_USDT")
	defer a.UnsubscribeFromChannel("BTC_USDT")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	ch := make(chan BestOrderBook)
	go a.ReadMessagesFromChannel(ch)
	select {
	case <-ch:
		return
	case <-ctx.Done():
		t.FailNow()
	}
}

