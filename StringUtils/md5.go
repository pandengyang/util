package StringUtils

import (
	"crypto/md5"
	"fmt"
)

func Md5(content string) string {
	md5 := md5.New()
	md5.Write([]byte(content))
	checksum := md5.Sum(nil)

	return fmt.Sprintf("%X", checksum)
}
