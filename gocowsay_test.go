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
	got := Inflator(inputList, maxWidth)
	want := []string{inputList[0] + strings.Repeat(" ", maxWidth-len(inputList[0])), inputList[1] + strings.Repeat(" ", maxWidth-len(inputList[1]))}

	assertErrors(t, got, want)
}

func TestSanitizeSpaces(t *testing.T) {
	inputString := `Lorem ipsum dolor sit amet, consectetur
elit, sed do eiusmod tempor incididunt ut labore
dolore magna aliqua. Ut enim ad minim veniam,
`
	got := SanitizeSpaces(inputString)
	want := "Lorem ipsum dolor sit amet, consectetur elit, sed do eiusmod tempor incididunt ut labore dolore magna aliqua. Ut enim ad minim veniam, "

	if got != want {
		t.Errorf("Got this string \"%s\" but want \"%s\"", got, want)
	}
}

func assertErrors(t testing.TB, got, want []string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v but want %v", got, want)
	}
}
