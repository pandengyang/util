package StringUtils

import (
	"crypto/sha256"
	"fmt"
)

func Sha256PasswdSalt(passwd, salt string) string {
	s256 := sha256.New()

	s256.Write([]byte(passwd))
	sum := s256.Sum([]byte(salt))

	hashedSaltedPasswd := fmt.Sprintf("%X", sum)

	return hashedSaltedPasswd
}
