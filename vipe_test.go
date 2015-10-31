package govipe

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockEditor struct {
	file      *os.File
	processor func(f *os.File)
}

func (m *mockEditor) Edit(filename string) error {
	f, err := os.OpenFile(filename, os.O_RDWR, 0600)
	m.file = f
	if m.processor != nil {
		m.file.Seek(0, 2)
		m.processor(m.file)
	}
	return err
}

func TestTemporaryFile(t *testing.T) {
	mock := &mockEditor{}
	SetEditor(mock)

	_, errVipe := Vipe(strings.NewReader("hello world!"))
	assert.Nil(t, errVipe)

	content, errRead := ioutil.ReadAll(mock.file)
	assert.Nil(t, errRead)
	assert.Equal(t, "hello world!", string(content))

	info, errStat := mock.file.Stat()
	assert.Nil(t, errStat)
	assert.Equal(t, uint32(0600), uint32(info.Mode().Perm()))
}

func TestOutput(t *testing.T) {
	mock := &mockEditor{}
	mock.processor = func(f *os.File) {
		fmt.Fprint(f, " world!")
	}
	SetEditor(mock)

	outReader, errVipe := Vipe(strings.NewReader("hello"))
	assert.Nil(t, errVipe)

	content, errRead := ioutil.ReadAll(outReader)
	log.Println(string(content))
	assert.Nil(t, errRead)
	assert.Equal(t, "hello world!", string(content))
}
