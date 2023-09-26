package leader

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/creack/pty"
)

type (
	Leader struct {
		file *os.File
	}
	shell string
)

const (
	Zsh  shell = "zsh"
	Bash shell = "bash"
	Fish shell = "fish"
)

func New() *Leader {
	return &Leader{}
}

func (l *Leader) Initialize(shell shell) error {
	c := exec.Command(string(shell))

	file, err := pty.Start(c)
	if err != nil {
		return errors.New("could not start shell")
	}
	l.file = file

	// Handle pty size.
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGWINCH)
	go func() {
		for range ch {
			if err := pty.InheritSize(os.Stdin, l.file); err != nil {
				log.Printf("error resizing pty: %s", err)
			}
		}
	}()
	ch <- syscall.SIGWINCH // Initial resize.
	//defer func() { signal.Stop(ch); close(ch) }() // Cleanup signals when done.

	return nil
}

func (l *Leader) Write(message string) (int, error) {
	return l.file.WriteString(message)
}

func (l *Leader) Read() (string, error) {
	arr := make([]byte, 16384)
	l.file.Read(arr)
	return string(arr), nil
}

func (l *Leader) Close() error {
	return l.file.Close()
}
