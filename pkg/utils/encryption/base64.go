package encryption

import "encoding/base64"

func EncryptBase64(value string) string {
	return base64.StdEncoding.EncodeToString([]byte(value))
}

func DecryptBase64(b64 string) (string, error) {
	res, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return "", err
	}
	return string(res), nil
}
