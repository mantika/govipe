package govipe

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Runnable interface {
}

func TestRunEditorSetInEnvironment(t *testing.T) {
	os.Setenv("EDITOR", "vim")

	editor := &systemDefaultEditor{}
	editor.Runner = func(cmd *exec.Cmd) error {
		assert.Equal(t, []string{"vim", "something.txt"}, cmd.Args)

		return nil
	}
	editor.Edit("something.txt")
}

func TestErrorIfVariableUnset(t *testing.T) {
	os.Setenv("EDITOR", "")
	editor := newSystemDefaultEditor()
	err := editor.Edit("something.txt")
	if err.Error() != "EDITOR variable not set" {
		t.Fail()
	}
}
