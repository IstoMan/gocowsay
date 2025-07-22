package main

import (
	"fmt"
	"io"
	"os"
)

func Greet(w io.Writer, text string) {
	fmt.Fprint(w, text)
}

func main() {
	Greet(os.Stdout, "Hello, World\n")
}
