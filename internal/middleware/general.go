package middleware

import (

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	cg "warabiz/api/pkg/constants/general"
)
 
func (mw *MiddlewareManager) RequestID() fiber.Handler {
	return requestid.New(requestid.ConfigDefault)
}

func (mw *MiddlewareManager) CORS() fiber.Handler {

	//* CORS
	CORS := cors.Config{}
	CORS.AllowHeaders = mw.cfg.Routes.Headers
	CORS.AllowMethods = mw.cfg.Routes.Methods
	if mw.cfg.Routes.DisableOrigins {
		CORS.AllowOrigins = "*"
	} else {
		CORS.AllowCredentials = true
		CORS.AllowOrigins = mw.cfg.Routes.Origins
	}

	return cors.New(CORS)
}

func (mw *MiddlewareManager) TimeZoneHeader() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tz := c.Get("TimeZone")
		if tz != "" {
			c.Locals(cg.CtxTimeZone)
		}
		return c.Next()
	}
}

func (mw *MiddlewareManager) Recover() fiber.Handler {
	return recover.New(recover.ConfigDefault)
}

// func (mw *MiddlewareManager) BearerToken() fiber.Handler {
// 	return func(c *fiber.Ctx) error {

// 		var token string
// 		var err error

// 		exc := exception.NewException(c, mw.logger)

// 		//* Get token
// 		token = getTokenCookie(c, mw.cfg.Cookies.CoreRT)
// 		if token == "" {
// 			token, err = getTokenHeader(c)
// 			if err != nil {
// 				return exc.WriteErrorResponse(http.StatusUnauthorized, "format token tidak valid", err)
// 			}
// 		}

// 		c.Locals("BearerToken", token)

// 		return c.Next()
// 	}
// }
