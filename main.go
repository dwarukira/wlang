package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/dwarukira/wakanda/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s Welcome to Wakanda 0.1.1 \n", user.Username)
	repl.Start(os.Stdin, os.Stdout)

}
