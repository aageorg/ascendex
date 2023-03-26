package main

import (
	"context"
	"encoding/json"
	"testing"
	"time"
)

type ReadClientOk struct{}

func (cl ReadClientOk) Close() error {
	return nil
}

func (cl ReadClientOk) ReadJSON(v interface{}) error {
	json.Unmarshal([]byte(`{"m":"bbo","symbol":"BTC/USDT","data":{"ts":1573068442532,"bid":["9309.11","0.0197172"],"ask":["9309.12","0.8851266"]}}`), v)
	return nil
}

func (cl ReadClientOk) WriteJSON(v interface{}) error {
	return nil
}

func TestReadMessagesFromChannel_Success(t *testing.T) {
	a := &Ascendex{
		conn: ReadClientOk{},
	}
	a.SubscribeToChannel("BTC_USDT")
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

type ReadClientBadData struct{}

func (cl ReadClientBadData) Close() error {
	return nil
}

func (cl ReadClientBadData) ReadJSON(v interface{}) error {
	json.Unmarshal([]byte(`{"m":"bbo","symbol":"BTC/USDT","data":{"ts":1573068442532,"bid":["93A9.11","0.01L7172"],"ask":["9309.12","0.8851266"]}}`), v)
	return nil
}

func (cl ReadClientBadData) WriteJSON(v interface{}) error {
	return nil
}

func TestReadMessagesFromChannel_BadData(t *testing.T) {
	a := &Ascendex{
		conn: ReadClientBadData{},
	}
	a.SubscribeToChannel("BTC_USDT")
	defer a.UnsubscribeFromChannel("BTC_USDT")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	ch := make(chan BestOrderBook)
	go a.ReadMessagesFromChannel(ch)
	select {
	case <-ch:
		t.FailNow()
	case <-ctx.Done():
		return
	}
}

type ReadClientNoMessage struct{}

func (cl ReadClientNoMessage) Close() error {
	return nil
}

func (cl ReadClientNoMessage) ReadJSON(v interface{}) error {
	return nil
}

func (cl ReadClientNoMessage) WriteJSON(v interface{}) error {
	return nil
}

func TestReadMessagesFromChannel_NoMessage(t *testing.T) {
	a := &Ascendex{
		conn: ReadClientNoMessage{},
	}
	a.SubscribeToChannel("BTC_USDT")
	defer a.UnsubscribeFromChannel("BTC_USDT")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	ch := make(chan BestOrderBook)
	go a.ReadMessagesFromChannel(ch)
	select {
	case <-ch:
		t.FailNow()
	case <-ctx.Done():
		return
	}
}
