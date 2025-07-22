package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	inputString := "Hello, Suhas"
	bytes := bytes.Buffer{}

	Greet(&bytes, inputString)

	got := bytes.String()
	want := "Hello, Suhas"

	if got != want {
		t.Errorf("got %s but wanted %s", got, want)
	}
}
