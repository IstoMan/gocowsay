package main

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

func TestGreet(t *testing.T) {
	inputString := "Hello, Suhas"
	textBuffer := bytes.Buffer{}

	Greet(&textBuffer, inputString)

	got := textBuffer.String()
	want := "Hello, Suhas"

	if got != want {
		t.Errorf("got %s but wanted %s", got, want)
	}
}

func TestSpilter(t *testing.T) {
	inputString := "Hello, World"

	got := Spilter(inputString)
	want := strings.Fields(inputString)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("the number of words got was %v but want %v", got, want)
	}
}
