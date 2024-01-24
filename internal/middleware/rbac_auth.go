package middleware

// import (
// 	"context"
// 	"errors"
// 	"fmt"
// 	"net/http"
// 	"strings"
// 	"time"

// 	"github.com/gofiber/fiber/v2"
// 	"golang.org/x/exp/slices"

// 	"project-pln/internal/models/session"
// 	"project-pln/pkg/http/exception"
// 	"project-pln/pkg/infra/db"
// 	"project-pln/pkg/utils/auth"
// 	"project-pln/pkg/utils/encryption"
// )

// // ===================================================================================

// func (mw *MiddlewareManager) RBAC() fiber.Handler {
// 	return func(c *fiber.Ctx) error {

// 		var err error
// 		var token string
// 		exc := exception.NewException(c, mw.logger)

// 		//* Get token
// 		token = getTokenCookie(c, mw.cfg.Cookies.CoreAT)
// 		if token == "" {
// 			token, err = getTokenHeader(c)
// 			if err != nil {
// 				return exc.WriteErrorResponse(http.StatusUnauthorized, "format token tidak valid", err.Error())
// 			}
// 		}

// 		//* Check EXP
// 		claims, err := auth.CheckJWTAccessToken(mw.cfg, token)
// 		if err != nil {
// 			return exc.WriteErrorResponse(http.StatusUnauthorized, "token kadaluwarsa", err.Error())
// 		}

// 		// * Check validation on cache
// 		sessData, err := mw.accessTokenValidation(c.Context(), exc, claims.Data.UserID, token)
// 		if err != nil {
// 			return exc.WriteParseError(err)
// 		}

// 		var sliceOfRoleID []int64
// 		for _, attr := range sessData.Role {
// 			sliceOfRoleID = append(sliceOfRoleID, attr.RoleID)
// 		}

// 		//* Get Path
// 		path := c.Route().Path
// 		pathPart := strings.Split(path, "/")
// 		lastPath := pathPart[len(pathPart)-1]
// 		if strings.Contains(lastPath, ":") {
// 			path = strings.ReplaceAll(path, "/"+lastPath, "")
// 		}

// 		err = mw.routeValidation(c.Context(), exc, route{
// 			Method: c.Method(),
// 			Path:   path,
// 		}, claims.Data.UserID, sliceOfRoleID)
// 		if err != nil {
// 			return exc.WriteParseError(err)
// 		}

// 		//* Set Active User Data
// 		c.Locals("active_user", sessData)
// 		return c.Next()
// 	}
// }

// func (mw *MiddlewareManager) RenewToken() fiber.Handler {
// 	return func(c *fiber.Ctx) error {

// 		var err error
// 		var token string
// 		exc := exception.NewException(c, mw.logger)

// 		//* Get token
// 		token = getTokenCookie(c, mw.cfg.Cookies.CoreRT)
// 		if token == "" {
// 			token, err = getTokenHeader(c)
// 			if err != nil {
// 				return exc.WriteErrorResponse(http.StatusUnauthorized, "format token tidak valid", err)
// 			}
// 		}

// 		//* Check EXP
// 		claims, err := auth.CheckJWTRefreshToken(mw.cfg, token)
// 		if err != nil {
// 			return exc.WriteErrorResponse(http.StatusUnauthorized, "token kadaluwarsa", err)
// 		}

// 		// * Check validation on cache
// 		sessData, err := mw.refreshTokenValidation(c.Context(), exc, claims.Data.UserID, token)
// 		if err != nil {
// 			return exc.WriteParseError(err)
// 		}

// 		//* Set Active User Data
// 		c.Locals("active_user", sessData)
// 		c.Locals("refresh_token", token)

// 		return c.Next()
// 	}
// }

// type route struct {
// 	Method string
// 	Path   string
// }

// func (mw MiddlewareManager) routeValidation(ctx context.Context, exc exception.Exception, route route, userID int64, sliceOfRoleID []int64) error {

// 	unique := encryption.GenerateSHA256(fmt.Sprintf("%s:%s", route.Method, route.Path))

// 	comparedData, err := mw.rmRedisRepo.GetRolesByUnique(ctx, unique)
// 	if err != nil {
// 		return exc.NewRestError(http.StatusInternalServerError, "failed to get rbac data", err.Error())
// 	}
// 	if comparedData != nil {
// 		ok := false
// 		for _, id := range sliceOfRoleID {
// 			if slices.Contains(*comparedData, id) {
// 				ok = true
// 				continue
// 			}
// 		}
// 		if !ok {
// 			return exc.NewRestError(http.StatusUnauthorized, "tidak ada izin akses", nil)
// 		}
// 	} else {

// 		//* DB Selector
// 		selectedDB, err := db.DBSelector(mw.dbList, mw.cfg.Connection.PLN.DriverSource)
// 		if err != nil {
// 			return exc.NewRestError(http.StatusInternalServerError, "unkown database address", err.Error())
// 		}

// 		params := make([]interface{}, 0)
// 		params = append(params, route.Method, route.Path)
// 		roleID, err := mw.rmRepo.GetAllRoleIDByRoute(ctx, selectedDB, params...)
// 		if err != nil {
// 			return exc.NewRestError(http.StatusInternalServerError, "failed to authenticate role access", err.Error())
// 		}

// 		ok := false
// 		for _, id := range sliceOfRoleID {
// 			if slices.Contains(roleID, id) {
// 				ok = true
// 				continue
// 			}
// 		}
// 		if !ok {
// 			return exc.NewRestError(http.StatusForbidden, "tidak ada izin akses", nil)
// 		} else {
// 			if err := mw.rmRedisRepo.StoreRoles(ctx, sliceOfRoleID, unique, mw.cfg.Authorization.JWT.AccessTokenDuration*time.Minute); err != nil {
// 				return exc.NewRestError(http.StatusInternalServerError, "failed to cache rbac data", err.Error())
// 			}
// 		}
// 	}
// 	return nil
// }

// func (mw MiddlewareManager) accessTokenValidation(ctx context.Context, exc exception.Exception, userID int64, accessToken string) (*session.UserSession, error) {
// 	sessData, err := mw.sessRedisRepo.GetSessionByID(ctx, userID)
// 	if err != nil {
// 		return nil, exc.NewRestError(http.StatusInternalServerError, "failed to get session data", err.Error())
// 	}
// 	if sessData == nil {
// 		return nil, exc.NewRestError(http.StatusBadRequest, "account anda telah logout, silahkan login kembali", nil)
// 	}
// 	// if sessData.AccessToken != accessToken {
// 	// 	return nil, exc.NewRestError(http.StatusConflict, "token tidak valid atau akun sedang login di tempat lain, silahkan login kembali", nil)
// 	// }
// 	return sessData, nil
// }

// func (mw MiddlewareManager) refreshTokenValidation(ctx context.Context, exc exception.Exception, userID int64, refreshToken string) (*session.UserSession, error) {
// 	sessData, err := mw.sessRedisRepo.GetSessionByID(ctx, userID)
// 	if err != nil {
// 		return nil, exc.NewRestError(http.StatusInternalServerError, "failed to get session data", err.Error())
// 	}
// 	if sessData == nil {
// 		return nil, exc.NewRestError(http.StatusBadRequest, "account anda telah logout, silahkan login kembali", nil)
// 	}
// 	// if sessData.RefreshToken != refreshToken {
// 	// 	return nil, exc.NewRestError(http.StatusConflict, "token tidak valid atau akun sedang login di tempat lain, silahkan login kembali", nil)
// 	// }
// 	return sessData, nil
// }

// func getTokenCookie(c *fiber.Ctx, key string) string {
// 	return c.Cookies(key)
// }

// func getTokenHeader(c *fiber.Ctx) (string, error) {
// 	authorizationHeader := c.Get("Authorization")
// 	if authorizationHeader == "" {
// 		return "", errors.New("token not found")
// 	}
// 	if !strings.Contains(authorizationHeader, "Bearer") {
// 		return "", errors.New("get token header error")
// 	}
// 	return strings.Replace(authorizationHeader, "Bearer ", "", -1), nil
// }
