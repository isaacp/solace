package main

// A simple program demonstrating the text input component from the Bubbles
// component library.

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/creack/pty"
	"github.com/solace/terminal"
)

var p *os.File
var term terminal.Terminal

type leader struct {
	file *os.File
}

func (l *leader) Initialize() error {
	c := exec.Command("/bin/zsh")
	file, err := pty.Start(c)
	if err != nil {
		return errors.New("could not start shell")
	}
	l.file = file
	return nil
}

func (l *leader) Write(message string) (int, error) {
	return l.file.WriteString(message)
}

func (l *leader) Read() (string, error) {
	arr := make([]byte, 16384)
	l.file.Read(arr)
	return string(arr), nil
}

func (l *leader) Close() error {
	return l.file.Close()
}

var l leader

func main() {
	// c := exec.Command("/bin/zsh")
	// pt, err := pty.Start(c)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// p = pt

	l = leader{}
	l.Initialize()
	defer l.Close()
	//term.Run()
	pr := tea.NewProgram(initialModel())
	if _, err := pr.Run(); err != nil {
		log.Fatal(err)
	}
}

type (
	errMsg error
)

type model struct {
	textInput textinput.Model
	err       error
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "<command>"
	ti.Focus()
	ti.CharLimit = 300
	ti.Width = 300

	return model{
		textInput: ti,
		err:       nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			l.Write(m.textInput.Value() + "\r")
			time.Sleep(100 * time.Millisecond)
			// t := make([]byte, 16092)
			t, _ := l.Read()
			fmt.Println(t)
			m.textInput.Reset()
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf("%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/solace/terminal"
// )

// func main() {
// 	term := terminal.NewTerminal(terminal.Zsh)
// 	term.OpenShell()
// 	if err := term.Run(); err != nil {
// 		log.Fatal(err)
// 	}

// 	line := ""
// 	for line != "quit" {
// 		fmt.Scanln(&line)
// 		term.Write(line)
// 		time.Sleep(1000)
// 		message, _ := term.Read()
// 		fmt.Println(message)
// 	}
// }
