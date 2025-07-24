package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

var errEmptyString = errors.New("cannot take empty string")

const maxWidth = 39

func Inflator(lines []string) []string {
	var ret []string
	for _, l := range lines {
		s := l + strings.Repeat(" ", maxWidth-utf8.RuneCountInString(l))
		ret = append(ret, s)
	}
	return ret
}

func Spiltter(text string) []string {
	words := strings.FieldsSeq(text)
	lines := []string{}
	var currentline string

	for word := range words {
		if len(currentline)+len(word)+1 > maxWidth {
			lines = append(lines, currentline)
			currentline = word
		} else {
			if len(currentline) != 0 {
				currentline += " " + word
			} else {
				currentline = word
			}
		}
	}

	if len(currentline) != 0 {
		lines = append(lines, currentline)
	}
	return lines
}

func DialogueBox(lines []string) []string {
	var ret []string
	top := strings.Repeat("-", maxWidth)
	ret = append(ret, top)
	count := len(lines)

	borders := [5]string{"/", "\\", "<", ">", "|"}

	for i, line := range lines {
		if count > 1 {
			if i == 0 {
				s := borders[0] + line + borders[1]
				ret = append(ret, s)
			}
			s := borders[4] + line + borders[4]
			ret = append(ret, s)
			if i == len(lines)-1 {
				s := borders[1] + line + borders[0]
				ret = append(ret, s)
			}
		}
	}

	ret = append(ret, top)

	return ret
}

func main() {
	inputString := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do"
	lines := Spiltter(inputString)
	padded := Inflator(lines)
	boxed := DialogueBox(padded)
	for _, line := range boxed {
		fmt.Println(line)
	}
}
