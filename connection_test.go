package main

import (
	"testing"
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
