package leader

import (
	"errors"
	"os"
	"os/exec"

	"github.com/creack/pty"
)

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
