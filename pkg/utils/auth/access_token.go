package auth

//! Don't delete!!
// clientSecret := utils.GenerateSHA256(fmt.Sprintf("%s:%s:%v", clientApiKey, clientApiSecret, requestTime))
// clientSecretForToken := utils.GenerateSHA256(fmt.Sprintf("%s:%s:%v", clientApiKey, clientSecret, ticks))

// func GenerateAccessToken(cfg *config.Config, clientID string, clientApiKey string, clientApiSecret string, ticks int) (string, error) {
// 	var clientIP string
// 	hash := fmt.Sprintf("%s:%s:%s:%v", clientIP, clientApiKey, "", ticks)
// 	h := hmac.New(sha256.New, []byte(clientApiSecret))
// 	h.Write([]byte(hash))
// 	hashLeft := hex.EncodeToString(h.Sum(nil))
// 	hashRight := fmt.Sprintf("%s:%s:%s:%v", clientID, clientApiKey, clientApiSecret, ticks)
// 	token, err := encryption.GetEncrypt([]byte(cfg.Server.HexaSecretKey), base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%v:%v", hashLeft, hashRight))))
// 	if err != nil {
// 		return "", err
// 	}
// 	return token, nil
// }

// func OpenAccessToken(token string, clientApiKey string) (string, error) {
// 	b64, err := encryption.GetDecrypt([]byte(clientApiKey), token)
// 	if err != nil {
// 		return "", err
// 	}
// 	bytes, err := base64.StdEncoding.DecodeString(b64)
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(bytes), nil
// }
