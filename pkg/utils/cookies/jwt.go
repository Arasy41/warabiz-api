package cookies

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"warabiz/api/config"
	// "warabiz/api/internal/models/user"
)

// func SetJWTCookie(c *fiber.Ctx, cfg *config.Config, token user.Token) {
// 	c.Cookie(&fiber.Cookie{
// 		Name:     cfg.Cookies.CoreAT,
// 		Value:    token.AccessToken,
// 		HTTPOnly: true,
// 		Secure:   true,
// 		Expires:  token.ATExp,
// 		Domain:   cfg.Cookies.CoreDomain,
// 		SameSite: "none",
// 	})
// 	c.Cookie(&fiber.Cookie{
// 		Name:     cfg.Cookies.CoreRT,
// 		Value:    token.RefreshToken,
// 		HTTPOnly: true,
// 		Secure:   true,
// 		Expires:  *token.RTExp,
// 		Domain:   cfg.Cookies.CoreDomain,
// 		SameSite: "none",
// 	})
// }

// func RenewJWTCookie(c *fiber.Ctx, cfg *config.Config, token user.Token) {
// 	c.Cookie(&fiber.Cookie{
// 		Name:     cfg.Cookies.CoreAT,
// 		Value:    token.AccessToken,
// 		HTTPOnly: true,
// 		Secure:   true,
// 		Expires:  token.ATExp,
// 		Domain:   cfg.Cookies.CoreDomain,
// 		SameSite: "none",
// 	})
// }

func ClearJWTCookie(c *fiber.Ctx, cfg *config.Config) {
	c.Cookie(&fiber.Cookie{
		Name:     cfg.Cookies.CoreAT,
		Value:    "",
		MaxAge:   -1,
		HTTPOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(-100 * time.Hour),
		Domain:   cfg.Cookies.CoreDomain,
		SameSite: "none",
	})

	c.Cookie(&fiber.Cookie{
		Name:     cfg.Cookies.CoreRT,
		Value:    "",
		MaxAge:   -1,
		HTTPOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(-100 * time.Hour),
		Domain:   cfg.Cookies.CoreDomain,
		SameSite: "none",
	})
}
