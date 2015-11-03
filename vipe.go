package govipe

import (
	"io"
	"io/ioutil"
	"os"
)

// Editor is the interface used by different supported editors
type Editor interface {
	Edit(file string) error
}

var setEditor Editor

// SetEditor sets the default editor
func SetEditor(editor Editor) {
	setEditor = editor
}

// GetEditor returns the current editor
func GetEditor() Editor {
	if setEditor == nil {
		setEditor = newSystemDefaultEditor()
	}
	return setEditor
}

const tempFilePrefix = "govipe"

// Vipe receives a Reader and sends its contect
// to the default editor
func Vipe(input io.Reader) (io.Reader, error) {
	file, errFile := ioutil.TempFile(os.TempDir(), tempFilePrefix)
	if errFile != nil {
		return nil, errFile
	}
	defer os.Remove(file.Name())

	_, errCopy := io.Copy(file, input)
	if errCopy != nil {
		return nil, errCopy
	}

	editor := GetEditor()
	if err := editor.Edit(file.Name()); err != nil {
		return nil, err
	}
	_, errSeek := file.Seek(0, 0)
	if errSeek != nil {
		file.Close()
		var err error
		if file, err = os.Open(file.Name()); err != nil {
			return nil, err
		}
	}
	return file, nil
}
