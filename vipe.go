package govipe

import (
	"io"
	"io/ioutil"
	"os"
)

type Editor interface {
	Edit(file string) error
}

var setEditor Editor

func SetEditor(editor Editor) {
	setEditor = editor
}
func GetEditor() Editor {
	if setEditor == nil {
		setEditor = NewSystemDefaultEditor()
	}
	return setEditor
}

const tempFilePrefix = "govipe"

func Vipe(input io.Reader) (io.Reader, error) {
	file, errFile := ioutil.TempFile(os.TempDir(), tempFilePrefix)
	if errFile != nil {
		return nil, errFile
	}
	file.Chmod(os.ModeTemporary | 0600)
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
		if file, err := os.Open(file.Name()); err != nil {
			return nil, err
		} else {
			return file, nil
		}
	}
	return file, nil
}
