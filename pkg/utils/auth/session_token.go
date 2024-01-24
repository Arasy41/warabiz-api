package auth

// func GenerateSessionToken(cfg *config.Config, userID string, clientApiKey string, channel string, ticks int) (string, error) {
// 	token, err := encryption.GetEncrypt([]byte(cfg.Server.HexaSecretKey), base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s:%s:%v", userID, clientApiKey, channel, ticks))))
// 	if err != nil {
// 		return "", err
// 	}
// 	return token, nil
// }

// func DecryptSessionToken(cfg *config.Config, token string) (int64, string, string, int, error) {
// 	dec, err := encryption.GetDecrypt([]byte(cfg.Server.HexaSecretKey), token)
// 	if err != nil {
// 		return 0, "", "", 0, err
// 	}
// 	dec64, err := base64.StdEncoding.DecodeString(dec)
// 	if err != nil {
// 		return 0, "", "", 0, err
// 	}
// 	parts := strings.Split(string(dec64), ":")
// 	if len(parts) != 4 {
// 		return 0, "", "", 0, errors.New("Invalid format token")
// 	}
// 	userId, err := strconv.Atoi(parts[0])
// 	if err != nil {
// 		return 0, "", "", 0, errors.New("Invalid format userID")
// 	}
// 	clientApiKey := parts[1]
// 	channel := parts[2]
// 	ticks, err := strconv.Atoi(parts[3])
// 	if err != nil {
// 		return 0, "", "", 0, errors.New("Invalid format ticks")
// 	}
// 	return int64(userId), clientApiKey, channel, ticks, err
// }
