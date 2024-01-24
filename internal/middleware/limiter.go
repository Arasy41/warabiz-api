package middleware

import (
	"fmt"
	"net/http"
	"warabiz/api/pkg/http/exception"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func (mw *MiddlewareManager) Limitter() fiber.Handler {

	return limiter.New(limiter.Config{
		Max:        1000,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			exc := exception.NewException(c, mw.logger)
			return exc.WriteErrorResponse(http.StatusBadRequest, fmt.Sprintf("You have requested too many in a single time-frame! Please wait another minute!"), nil)
		},
	})
}

//* Limiter middleware
func (mw *MiddlewareManager) AddLimit(max int, exp int, skipFailed bool, skipSuccess bool) fiber.Handler {
	return limiter.New(
		limiter.Config{
			Max:        max,
			Expiration: time.Duration(exp) * time.Minute,
			KeyGenerator: func(c *fiber.Ctx) string {
				return c.IP()
			},
			LimitReached: func(c *fiber.Ctx) error {
				exc := exception.NewException(c, mw.logger)
				return exc.WriteErrorResponse(http.StatusBadRequest, fmt.Sprintf("Terlalu banyak upaya login yang gagal, silahkan coba lagi %v menit kemudian", exp), nil)
			},
			SkipFailedRequests:     false,
			SkipSuccessfulRequests: true,
			LimiterMiddleware:      limiter.FixedWindow{},
		},
	)
}

// func LimiterChangePwd() {
// 	app := initData.App
// 	app.Use("/api/v1/user/changepassword", limiter.New(limiter.Config{
// 		Max:        5,
// 		Expiration: 1 * time.Minute,
// 		KeyGenerator: func(c *fiber.Ctx) string {
// 			return c.IP()
// 		},
// 		LimitReached: func(c *fiber.Ctx) error {
// 			init := exception.InitException(c, initData.Conf, initData.Log)

// 			return exception.CreateResponse_Log(init, http.StatusBadRequest, "Too many failed attempts, please try again 1 minute later ", "Terlalu banyak upaya yang gagal, silahkan coba lagi 1 menit kemudian", nil)
// 		},
// 		SkipFailedRequests:     false,
// 		SkipSuccessfulRequests: true,
// 		LimiterMiddleware:      limiter.FixedWindow{},
// 	}))
// }
