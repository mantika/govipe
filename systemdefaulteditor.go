package govipe

import (
	"errors"
	"os"
	"os/exec"
)

type systemDefaultEditor struct {
	Runner func(*exec.Cmd) error
}

func newSystemDefaultEditor() *systemDefaultEditor {
	return &systemDefaultEditor{
		Runner: func(cmd *exec.Cmd) error {
			return cmd.Run()
		},
	}
}

func (s *systemDefaultEditor) Edit(filename string) error {
	command := os.Getenv("EDITOR")

	if command == "" {
		errors.New("EDITOR variable not set")
	}

	cmd := exec.Command(command, filename)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	errCmd := s.Runner(cmd)

	return errCmd
}
