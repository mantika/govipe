package govipe

import (
	"io/ioutil"
	"os"
	"os/exec"
)

const tempFilePrefix = "govipe"

// Runner is used to execute the corresponding editor command.
// Eventually different platforms can implement it's own Runner.
type Runner interface {
	Run(command *exec.Cmd) error
}

// CommandRunner calls Run on it's supplied Cmd argument.
type CommandRunner struct{}

// Run performs the command execution to make the editing
func (e CommandRunner) Run(command *exec.Cmd) error {
	return command.Run()
}

var defaultRunner Runner = CommandRunner{}

// Edit opens the default editor with the specified input,
// and returns the modified output.
func Edit(input []byte) ([]byte, error) {
	file, errFile := ioutil.TempFile(os.TempDir(), tempFilePrefix)
	if errFile != nil {
		return nil, errFile
	}
	defer os.Remove(file.Name())

	fileMode := os.FileMode(0600)
	errWrite := ioutil.WriteFile(file.Name(), input, os.ModeTemporary|fileMode)
	if errWrite != nil {
		return nil, errWrite
	}

	cmd := getCommand(file.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	errCmd := defaultRunner.Run(cmd)

	if errCmd != nil {
		return nil, errCmd
	}

	out, errOut := ioutil.ReadFile(file.Name())
	if errOut != nil {
		return nil, errOut
	}

	return out, nil
}

func getCommand(filename string) *exec.Cmd {
	return exec.Command(os.Getenv("EDITOR"), filename)
}
