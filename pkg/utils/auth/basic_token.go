package auth

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"golang.org/x/exp/utf8string"
)

func GenerateBasicToken(apiKey, apiSecret, ticks string) string {
	message := fmt.Sprintf("%s:%s:%v", apiKey, apiSecret, ticks)
	hash := sha256.Sum256([]byte(message))
	wordArray := utf8string.NewString(fmt.Sprintf("%s:%x", apiKey, hash))
	secret_buffer := base64.StdEncoding.EncodeToString([]byte(wordArray.String()))
	return secret_buffer
}
