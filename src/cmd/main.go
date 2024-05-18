package main

import (
	"NoobAuthor/interpreter/src/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s!\n", user.Username)
	fmt.Printf("This is the Monkey programming language!\n")
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)

}
