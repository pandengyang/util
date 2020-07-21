package FileUtils

import (
	"os"
	"path/filepath"
	"time"
)

func MkdirAllByDate(parentPath string) (path string, err error) {
	path = filepath.Join(parentPath, time.Now().Format("2006-01-02"))
	err = os.MkdirAll(path, os.ModePerm)

	return path, err
}
