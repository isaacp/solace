package main

import (
	"fmt"
	"log"
	"time"

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
		term.Write(line)
		time.Sleep(1000)
		message, _ := term.Read()
		fmt.Println(message)
	}
}
