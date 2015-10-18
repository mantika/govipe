package govipe

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockRunner struct{}

func (e MockRunner) Run(command *exec.Cmd) error {
	f, _ := os.OpenFile(command.Args[1], os.O_APPEND|os.O_WRONLY, 0600)
	f.WriteString(" world!")
	return nil
}

func TestEdit(t *testing.T) {
	defaultRunner = MockRunner{}
	out, err := Edit([]byte("hello"))

	assert.Nil(t, err)
	assert.Equal(t, "hello world!", string(out))
}

func TestCommand(t *testing.T) {
	os.Setenv("EDITOR", "somecommand")
	cmd := getCommand("somefile")

	assert.Equal(t, []string{"somecommand", "somefile"}, cmd.Args)
}
