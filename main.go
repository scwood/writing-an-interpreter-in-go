package main

import (
	"os"

	"github.com/scwood/writing-an-interpreter-in-go/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
