package main

import (
	"errors"
	"testing"
)

type SubscribeClientOk struct{}

func (cl SubscribeClientOk) Close() error {
	return nil
}

func (cl SubscribeClientOk) ReadJSON(v interface{}) error {
	return nil
}

func (cl SubscribeClientOk) WriteJSON(v interface{}) error {
	return nil
}

func TestSubscribeToChannel_Success(t *testing.T) {
	a := &Ascendex{
		conn: SubscribeClientOk{},
	}
	err := a.SubscribeToChannel("BTC_USDT")
	defer a.UnsubscribeFromChannel("BTC_USDT")
	if err != nil {
		t.FailNow()
	}

}

func TestSubscribeToChannel_NilClient(t *testing.T) {
	a := &Ascendex{}
	err := a.SubscribeToChannel("BTC_USDT")
	if err == nil {
		t.FailNow()
	}
}

func TestSubscribeToChannel_BadParam(t *testing.T) {
	a := &Ascendex{
		conn: SubscribeClientOk{},
	}
	err := a.SubscribeToChannel("BTCUSDT")
	if err == nil {
		t.FailNow()
	}
}

type SubscribeClientFail struct{}

func (cl SubscribeClientFail) Close() error {
	return nil
}

func (cl SubscribeClientFail) ReadJSON(v interface{}) error {
	return nil
}

func (cl SubscribeClientFail) WriteJSON(v interface{}) error {
	return errors.New("Cannot write JSON")
}

func TestSubscribeToChannel_Failure(t *testing.T) {
	a := &Ascendex{
		conn: SubscribeClientFail{},
	}
	err := a.SubscribeToChannel("BTC_USDT")
	if err == nil {
		t.FailNow()
	}
}
