package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"

	"speak-monke/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	splitUsername := strings.Split(user.Username, "\\")
	fmt.Printf(
		"Hello %s! Welcome to the Monkey Programming Language! 🐒\n",
		splitUsername[len(splitUsername)-1],
	)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
