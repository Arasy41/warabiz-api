package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"net/url"
)

func EncryptAESGCM(message string, secretKey string) (string, error) {
	iv, enc, err := encryptAESGCM(message, secretKey)
	if err != nil {
		return "", err
	}
	result := url.PathEscape(iv + enc)
	return result, nil
}

func DecryptAESGCM(encryptedData string, secretKey string) (string, error) {
	str, err := url.PathUnescape(encryptedData)
	if err != nil {
		return "", err
	}
	if len(str) < 16 {
		return "", errors.New("invalid format")
	}
	dec, err := decryptAESGCM(str[16:], str[0:16], secretKey)
	if err != nil {
		return "", err
	}
	return dec, nil
}

func encryptAESGCM(message string, secretKey string) (string, string, error) {

	byteKey := parseSecretKey(secretKey)

	block, err := aes.NewCipher(byteKey)
	if err != nil {
		return "", "", err
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", "", err
	}

	encrypted := aesGCM.Seal(nil, nonce, []byte(message), nil)

	ivStr := base64.StdEncoding.EncodeToString(nonce)
	encryptedDataStr := base64.StdEncoding.EncodeToString(encrypted)

	return ivStr, encryptedDataStr, nil
}

func decryptAESGCM(encryptedData string, iv string, secretKey string) (string, error) {

	byteKey := parseSecretKey(secretKey)

	decodedEncryptedData, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return "", err
	}

	decodedIV, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(byteKey)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	plaintext, err := aesGCM.Open(nil, decodedIV, decodedEncryptedData, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func parseSecretKey(secretKey string) []byte {
	for len(secretKey) < 32 {
		secretKey += " "
	}
	return []byte(secretKey)
}
