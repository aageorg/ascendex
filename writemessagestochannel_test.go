package main

import (
	"context"
	"encoding/json"
	"testing"
	"time"
)

type PingClientOk struct{}

func (cl PingClientOk) Close() error {
	return nil
}

func (cl PingClientOk) ReadJSON(v interface{}) error {
	err := json.Unmarshal([]byte(`{"m":"pong"}`), v)
	return err
}

func (cl PingClientOk) WriteJSON(v interface{}) error {
	return nil
}

func TestWriteMessagesToChannel_Success(t *testing.T) {
	a := &Ascendex{
		conn: PingClientOk{},
	}
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

type PingClientFail struct{}

func (cl PingClientFail) Close() error {
	return nil
}

func (cl PingClientFail) ReadJSON(v interface{}) error {
	return nil
}

func (cl PingClientFail) WriteJSON(v interface{}) error {
	return nil
}

func TestWriteMessagesToChannel_NoAnswer(t *testing.T) {
	a = &Ascendex{
		conn: PingClientFail{},
	}
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
		t.FailNow()
	case <-ctx.Done():
		return
	}
}
