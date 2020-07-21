package FileUtils

import (
	"os"
	"path/filepath"
	"time"
)

func MkdirAllByDate(parentDir string) (err error) {
	return os.MkdirAll(filepath.Join(parentDir, time.Now().Format("2006-01-02")), os.ModePerm)
}
