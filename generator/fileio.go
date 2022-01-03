package generator

import (
	"io"
	"os"

	"github.com/borgmon/tunaflow/assets"
	"github.com/pkg/errors"
)

func CopyFile(src, dst string) error {
	file, err := assets.AssetFiles.Open(src)
	if err != nil {
		return errors.WithStack(err)
	}

	sourceFileStat, err := file.Stat()
	if err != nil {
		return errors.WithStack(err)
	}

	if !sourceFileStat.Mode().IsRegular() {
		return errors.WithStack(err)
	}

	destination, err := os.Create(dst)
	if err != nil {
		return errors.WithStack(err)
	}
	defer destination.Close()
	_, err = io.Copy(destination, file)
	return errors.WithStack(err)
}
