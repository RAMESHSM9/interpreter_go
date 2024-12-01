package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/RAMESHSM9/interpreter_go/src/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s, Welcome to Parrot Programming language!", user.Username)
	fmt.Println("feel free to start typing in the commnads")
	repl.Start(os.Stdin, os.Stdout)
}
