package middleware

// import (
// 	"context"
// 	"errors"
// 	"strings"

// 	"github.com/gofiber/fiber/v2"

// 	"suretybond/config"

// 	sessUc "suretybond/internal/session/usecase"
// )

// // ===================================================================================

// // * Role
// // func (mw *MiddlewareManager) JWTIsUser() fiber.Handler {
// // 	return CheckJWT(mw)
// // }

// func IsAuthenticated(c *fiber.Ctx, sessUc sessUc.Usecase, userID int64, accessToken string, refreshToken string) error {

// 	// ctx := locals.CreateContext(c)

// 	tokens, err := sessUc.GetSessionByID(context.Background(), userID)
// 	if err != nil {
// 		return errors.New("tokens not found")
// 	}

// 	if tokens != nil {
// 		if accessToken != "" {
// 			if accessToken != tokens.AccessToken {
// 				return errors.New("unauthorized")
// 			}
// 		}
// 		if refreshToken != "" {
// 			if refreshToken != tokens.RefreshToken {
// 				return errors.New("unauthorized")
// 			}
// 		}
// 	} else {
// 		return errors.New("unauthorized")
// 	}

// 	return nil
// }

// // func CheckJWT(mw *MiddlewareManager) fiber.Handler {
// // 	return func(c *fiber.Ctx) error {

// // 		var err error
// // 		var token string
// // 		exc := exception.NewException(languages.ParseLanguage(c.Locals(cg.CtxLanguage)), mw.logger)

// // 		//* Get token
// // 		token = getTokenCookie(c, mw.cfg)
// // 		if token == "" {
// // 			token, err = getTokenHeader(c)
// // 			if err != nil {
// // 				return exc.WriteErrorResponse( http.StatusUnauthorized, exc.Message.FormatError("token"), err)
// // 			}
// // 		}

// // 		//* Check EXP
// // 		claims, err := auth.CheckJWTAccessToken(mw.cfg, token)
// // 		if err != nil {
// // 			return exc.WriteErrorResponse( http.StatusUnauthorized, exc.Message.ExpiredTokenError(), err)
// // 		}

// // 		// * Check validation on cache
// // 		err = IsAuthenticated(c, mw.sessUc, claims.UserID, token, "")
// // 		if err != nil {
// // 			return exc.WriteErrorResponse( http.StatusUnauthorized, exc.Message.InvalidTokenError(), err)
// // 		}

// // c.Locals("active_user", user.ActiveUser{
// // 	UserID:   claims.UserID,
// // 	Username: claims.Username,
// // 	FullName: claims.FullName,
// // })

// // 		return c.Next()
// // 	}
// // }

// // func (mw *MiddlewareManager) RenewJWT() fiber.Handler {
// // 	return func(c *fiber.Ctx) error {

// // 		var err error
// // 		var token string
// // 		exc := exception.NewException(languages.ParseLanguage(c.Locals(cg.CtxLanguage)), mw.logger)

// // 		//* Get token
// // 		token = getTokenCookie(c, mw.cfg)
// // 		if token == "" {
// // 			token, err = getTokenHeader(c)
// // 			if err != nil {
// // 				return exc.WriteErrorResponse( http.StatusUnauthorized, exc.Message.FormatError("token"), err)
// // 			}
// // 		}

// // 		//* Check EXP
// // 		claims, err := auth.CheckJWTRefreshToken(mw.cfg, token)
// // 		if err != nil {
// // 			return exc.WriteErrorResponse( http.StatusUnauthorized, exc.Message.ExpiredTokenError(), err)
// // 		}

// // 		// * Check validation on cache
// // 		err = IsAuthenticated(c, mw.sessUc, claims.UserID, "", token)
// // 		if err != nil {
// // 			return exc.WriteErrorResponse( http.StatusUnauthorized, exc.Message.InvalidTokenError(), err)
// // 		}

// // 		c.Locals("active_user", user.ActiveUser{
// // 			UserID:   claims.UserID,
// // 			Username: claims.Username,
// // 			FullName: claims.FullName,
// // 		})
// // 		c.Locals("refresh_token", token)

// // 		return c.Next()
// // 	}
// // }

// func getTokenCookie(c *fiber.Ctx, cfg *config.Config) string {
// 	return c.Cookies(cfg.Cookies.CoreAT)
// }

// func getTokenHeader(c *fiber.Ctx) (string, error) {
// 	authorizationHeader := c.Get("Authorization")
// 	if !strings.Contains(authorizationHeader, "Bearer") {
// 		return "", errors.New("get token header error")
// 	}
// 	return strings.Replace(authorizationHeader, "Bearer ", "", -1), nil
// }
