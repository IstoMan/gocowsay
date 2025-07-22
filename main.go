package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func Spilter(text string) []string {
	return strings.Fields(text)
}

func Greet(w io.Writer, text string) {
	fmt.Fprint(w, text)
}

func main() {
	Greet(os.Stdout, "Hello, World\n")
}
