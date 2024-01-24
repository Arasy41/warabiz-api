package encryption

import (
	"crypto/md5"
	"fmt"
)

func GenerateUnique(key string) (string, error) {
	hasher := md5.New()
	hasher.Write([]byte(key))
	hash := hasher.Sum(nil)
	return fmt.Sprintf("%x", hash), nil
}
