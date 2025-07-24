package main

import (
	"reflect"
	"strings"
	"testing"
)

var inputText = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed"

func TestSpiltter(t *testing.T) {
	got := Spiltter(inputText)
	want := []string{"Lorem ipsum dolor sit amet, consectetur", "adipiscing elit, sed"}

	assertErrors(t, got, want)
}

func TestInflator(t *testing.T) {
	inputList := Spiltter(inputText)
	got := Inflator(inputList)
	want := []string{inputList[0] + strings.Repeat(" ", maxWidth-len(inputList[0])), inputList[1] + strings.Repeat(" ", maxWidth-len(inputList[1]))}

	assertErrors(t, got, want)
}

func assertErrors(t testing.TB, got, want []string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v but want %v", got, want)
	}
}
