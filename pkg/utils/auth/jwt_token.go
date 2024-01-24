package auth

import (
	"errors"
	"fmt"
	"warabiz/api/config"
	// "warabiz/api/internal/models/user"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTData struct {
	UserID   int64           `json:"user_id"`
	FullName string          `json:"full_name"`
	// Role     []user.UserRole `json:"role"`
}

// type ActiveUserRole struct {
// 	RoleID   int64  `json:"role_id"`
// 	RoleName string `json:"role_name"`
// }

type JWTAccessTokenClaims struct {
	Data JWTData
	jwt.StandardClaims
}

type JWTRefreshTokenClaims struct {
	Data JWTData
	jwt.StandardClaims
}

type JWTResponse struct {
	AccessToken  string    `json:"access_token"`
	ATExp        time.Time `json:"at_exp"`
	RefreshToken string    `json:"refresh_token"`
	RTExp        time.Time `json:"rt_exp"`
}

func GenerateJWTAccessToken(cfg *config.Config, data JWTData, timeNow *time.Time) (string, *time.Time, error) {

	claims := JWTAccessTokenClaims{Data: data}
	exp := timeNow.Add(cfg.Authorization.JWT.AccessTokenDuration * time.Minute)

	claims.IssuedAt = timeNow.Unix()
	claims.ExpiresAt = exp.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(cfg.Authorization.JWT.AccessTokenSecretKey))
	if err != nil {
		return "", nil, err
	}
	return tokenString, &exp, nil
}

func GenerateJWTRefreshToken(cfg *config.Config, data JWTData, timeNow *time.Time) (string, *time.Time, error) {

	claims := JWTRefreshTokenClaims{Data: data}
	exp := timeNow.Add(cfg.Authorization.JWT.AccessTokenDuration * time.Hour * 24)

	claims.IssuedAt = timeNow.Unix()
	claims.ExpiresAt = exp.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS384, claims)
	tokenString, err := token.SignedString([]byte(cfg.Authorization.JWT.RefreshTokenSecretKey))
	if err != nil {
		return "", nil, err
	}
	return tokenString, &exp, nil
}

func CheckJWTAccessToken(cfg *config.Config, accessToken string) (*JWTAccessTokenClaims, error) {

	claims := &JWTAccessTokenClaims{}

	token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.Authorization.JWT.AccessTokenSecretKey), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, errors.New("invalid token signature")
		}
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token ")
	}

	claims, ok := token.Claims.(*JWTAccessTokenClaims)
	if !ok {
		return nil, fmt.Errorf("access token valid but couldn't parse claims")
	}

	return claims, nil
}

func CheckJWTRefreshToken(cfg *config.Config, refreshToken string) (*JWTRefreshTokenClaims, error) {

	claims := &JWTRefreshTokenClaims{}

	token, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.Authorization.JWT.RefreshTokenSecretKey), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, errors.New("invalid token signature")
		}
		return nil, err
	}

	claims, ok := token.Claims.(*JWTRefreshTokenClaims)
	if !ok {
		return nil, fmt.Errorf("access token valid but couldn't parse claims")
	}

	return claims, nil
}
