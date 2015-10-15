package govipe

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEdit(t *testing.T) {
	os.Setenv("EDITOR", `sed -i 's/hello/hello world!/g'`)
	out, err := Edit([]byte("hello"))

	assert.Nil(t, err)
	assert.Equal(t, "hello world!", string(out))
}
