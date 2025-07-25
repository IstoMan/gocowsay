package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

var errEmptyString = errors.New("No text given to say.")

const maxWidth = 39

func SanitizeSpaces(text string) string {
	ret := strings.ReplaceAll(text, "\n", " ")
	ret = strings.ReplaceAll(ret, "\t", " ")

	return ret
}

func Inflator(lines []string, maxwidth int) []string {
	var ret []string
	for _, l := range lines {
		s := l + strings.Repeat(" ", maxwidth-utf8.RuneCountInString(l))
		ret = append(ret, s)
	}
	return ret
}

func Spiltter(text string) []string {
	words := strings.FieldsSeq(text)
	lines := []string{}
	var currentline string

	for word := range words {
		if utf8.RuneCountInString(currentline)+utf8.RuneCountInString(word)+1 > maxWidth {
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
	retMaxWidth := 0
	for _, line := range lines {
		lineWidth := utf8.RuneCountInString(line)
		if lineWidth > retMaxWidth {
			retMaxWidth = lineWidth
		}
	}
	return retMaxWidth
}

func cowSelected(option string) string {
	var ret string
	switch option {
	case "cow":
		ret = `         \  ^__^
          \ (oo)\_______
	    (__)\       )\/\
	        ||----w |
	        ||     ||
		`
	case "tux":
		ret = ` 
   \
    \
        .--.
       |o o |
       |:_/ |
      //    \\ 
     (|     | )
    /'\\_   _/'\
    \\___)=(___/
		`
	}
	return ret
}

func main() {
	cowChar := flag.String("f", "cow", "Character to display. Use 'tux' for Tux.")
	flag.Parse()

	inputString := strings.Join(flag.Args(), " ")

	if strings.TrimSpace(inputString) == "" {
		fmt.Fprintln(os.Stderr, errEmptyString)
		fmt.Fprintln(os.Stderr, "Usage: gosay [flags] <text to say>")
		flag.PrintDefaults()
		os.Exit(1)
	}

	cow := cowSelected(*cowChar)

	sanitizedText := SanitizeSpaces(inputString)
	lines := Spiltter(sanitizedText)
	maxwidth := findMaxWidth(lines)
	padded := Inflator(lines, maxwidth)
	boxed := DialogueBox(padded, maxwidth)

	fmt.Println(boxed)
	fmt.Println(cow)
}
