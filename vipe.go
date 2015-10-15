package govipe

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

const TempFilePrefix = "govipe"

// Edit opens the default editor with the specified input,
// and returns the modified output.
func Edit(input []byte) ([]byte, error) {
	file, errFile := ioutil.TempFile(os.TempDir(), TempFilePrefix)
	if errFile != nil {
		return nil, errFile
	}
	defer os.Remove(file.Name())

	fileMode := os.FileMode(0600)
	errWrite := ioutil.WriteFile(file.Name(), input, os.ModeTemporary|fileMode)
	if errWrite != nil {
		return nil, errWrite
	}

	command := os.Getenv("EDITOR")

	cmd := exec.Command("sh", "-c", fmt.Sprintf("%s %s", command, file.Name()))

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	errCmd := cmd.Run()

	if errCmd != nil {
		return nil, errCmd
	}

	out, errOut := ioutil.ReadFile(file.Name())
	if errOut != nil {
		return nil, errOut
	}

	return out, nil
}
