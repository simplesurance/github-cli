package flag

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/pflag"
)

type FilePathFlag struct {
	path string
}

func (f *FilePathFlag) String() string {
	return f.path
}

func (f *FilePathFlag) Set(val string) error {
	if val == "" {
		return errors.New("empty string is not valid path")
	}

	f.path = val

	return nil
}

func (f *FilePathFlag) Type() string {
	return "<path>"
}

func (f *FilePathFlag) FileContent() ([]byte, error) {
	var file *os.File

	if f.path == "-" {
		file = os.Stdin
	} else {
		var err error

		file, err = os.Open(f.path)
		if err != nil {
			return nil, fmt.Errorf("opening file %s failed: %w", f.path, err)
		}
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("reading %s failed: %w", file.Name(), err)
	}

	_ = file.Close()

	return content, err
}

func (f *FilePathFlag) FileContentString() (string, error) {
	content, err := f.FileContent()
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func (f *FilePathFlag) Path() string {
	return f.path
}

var _ pflag.Value = &FilePathFlag{}
