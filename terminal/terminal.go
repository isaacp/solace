package terminal

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/creack/pty"
)

type (
	Terminal struct {
		ptmx    *os.File
		execCmd *exec.Cmd
	}

	shell string
)

const (
	Zsh  shell = "/bin/zsh"
	Bash shell = "bash"
	Fish shell = "fish"
)

func NewTerminal(shell shell) *Terminal {

	return &Terminal{
		execCmd: exec.Command(string(shell)),
	}
}

func (t *Terminal) OpenShell() error {
	// Start the command with a pty.
	ptmx, err := pty.Start(t.execCmd)
	if err != nil {
		return fmt.Errorf("start error: %s", err)
	}
	t.ptmx = ptmx
	return nil
}

func (t *Terminal) CloseShell() error {
	return t.ptmx.Close()
}

func (t *Terminal) Write(message string) (int, error) {
	return t.ptmx.Write([]byte(message))
}
func (t *Terminal) Read() (string, error) {
	message := make([]byte, 0)
	_, err := t.ptmx.Read(message)
	if err != nil {
		return "", fmt.Errorf("reading: %s", err)
	}
	return string(message), nil
}

// func (t *Terminal) Run() error {
// 	// Make sure to close the pty at the end.
// 	defer func() { _ = t.CloseShell() }() // Best effort.

// 	// Handle pty size.
// 	ch := make(chan os.Signal, 1)
// 	signal.Notify(ch, syscall.SIGWINCH)
// 	go func() {
// 		for range ch {
// 			if err := pty.InheritSize(os.Stdin, t.ptmx); err != nil {
// 				log.Printf("error resizing pty: %s", err)
// 			}
// 		}
// 	}()
// 	ch <- syscall.SIGWINCH                        // Initial resize.
// 	defer func() { signal.Stop(ch); close(ch) }() // Cleanup signals when done.

// 	// Copy stdin to the pty and the pty to stdout.
// 	// NOTE: The goroutine will keep reading until the next keystroke before returning.
// 	// go func() { _, _ = io.Copy(t.ptmx, os.Stdin) }()
// 	// go func() { _, _ = io.Copy(os.Stdout, t.ptmx) }()

// 	return nil
// }
