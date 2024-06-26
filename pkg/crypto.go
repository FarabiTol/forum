package pkg

import (
	"crypto/sha1"
	"fmt"
)

const (
	salt = "mgfd#g5"
)

func GetPasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
