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
	args := os.Args[1:]
	// fmt.Println(args[0])
	if len(args) >= 1 {
		ReadFile(args[0], os.Stdout)
		return
	}
	fmt.Printf("%s Welcome to Wakanda 0.1.1 \n", user.Username)
	repl.Start(os.Stdin, os.Stdout)

}
