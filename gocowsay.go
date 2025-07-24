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

	top := " " + strings.Repeat("_", maxWidth+2)
	bottom := " " + strings.Repeat("-", maxWidth+2)

	ret = append(ret, top)
	count := len(lines)

	borders := [5]string{"/", "\\", "<", ">", "|"}

	if count == 1 {
		s := fmt.Sprintf("%s %s %s", borders[2], lines[0], borders[3])
		ret = append(ret, s)
		return ret
	} else {
		s := fmt.Sprintf("%s %s %s", borders[0], lines[0], borders[1])
		ret = append(ret, s)

		i := 1

		for ; i < count-1; i++ {
			s := fmt.Sprintf("%s %s %s", borders[4], lines[i], borders[4])
			ret = append(ret, s)
		}

		s = fmt.Sprintf("%s %s %s", borders[1], lines[i], borders[0])
		ret = append(ret, s)
	}

	ret = append(ret, bottom)

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
