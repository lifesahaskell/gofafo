package main

import (
	"fafolang/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s!, You've started the Fafolang REPL!\n", user.Username)
	fmt.Printf("Test out some code or enter a command.\n")
	repl.Start(os.Stdin, os.Stdout)
}
