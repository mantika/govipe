package govipe

import (
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

	cmd := exec.Command(command, filename)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	errCmd := s.Runner(cmd)

	return errCmd
}
