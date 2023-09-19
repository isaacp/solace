package main

import (
	"fmt"
	"log"

	"github.com/solace/terminal"
)

func main() {
	term := terminal.NewTerminal(terminal.Zsh)
	term.OpenShell()
	if err := term.Run(); err != nil {
		log.Fatal(err)
	}

	line := ""
	for line != "quit" {
		fmt.Scanln(&line)
		fmt.Println(line)
	}
}
