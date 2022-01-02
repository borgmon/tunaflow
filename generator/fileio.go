package generator

import (
	"io"
	"os"

	"github.com/pkg/errors"
)

func CopyFile(src, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return errors.WithStack(err)
	}

	if !sourceFileStat.Mode().IsRegular() {
		return errors.WithStack(err)
	}

	source, err := os.Open(src)
	if err != nil {
		return errors.WithStack(err)
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return errors.WithStack(err)
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	return errors.WithStack(err)
}
