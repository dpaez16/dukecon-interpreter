package main

import (
	"dukecon/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the DukeCON programming language!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
