package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var errEmptyString = errors.New("cannot take empty string")

const maxWidth = 39

func Inflator(lines []string, maxwidth int) []string {
	var ret []string
	for _, l := range lines {
		s := l + strings.Repeat(" ", maxwidth-len(l))
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

func DialogueBox(lines []string, maxwidth int) string {
	var ret []string
	count := len(lines)

	top := " " + strings.Repeat("_", maxwidth+2)
	bottom := " " + strings.Repeat("-", maxwidth+2)

	ret = append(ret, top)

	borders := [5]string{"/", "\\", "<", ">", "|"}

	if count == 1 {
		s := fmt.Sprintf("%s %s %s", borders[2], lines[0], borders[3])
		ret = append(ret, s)
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

	return strings.Join(ret, "\n")
}

func findMaxWidth(lines []string) int {
	retMaxWidth := maxWidth
	for _, line := range lines {
		if len(line) < maxWidth {
			retMaxWidth = len(line)
		}
	}
	return retMaxWidth
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No text given")
		fmt.Println(`gosay "-----"`)
		return
	}

	inputString := strings.Join(os.Args[1:], " ")
	lines := Spiltter(inputString)
	maxwidth := findMaxWidth(lines)
	padded := Inflator(lines, maxwidth)
	boxed := DialogueBox(padded, maxwidth)

	cow := `         \  ^__^
          \ (oo)\_______
	    (__)\       )\/\
	        ||----w |
	        ||     ||
		`

	fmt.Println(boxed)
	fmt.Println(cow)
}
