package cookies

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"warabiz/api/config"
)

func SetForPasTokenCookie(c *fiber.Ctx, cfg *config.Config, token string, exp time.Time) {
	c.Cookie(&fiber.Cookie{
		Name:     "forpas-token",
		Value:    token,
		HTTPOnly: true,
		Secure:   true,
		Expires:  exp,
		Domain:   cfg.Cookies.CoreDomain,
		SameSite: "none",
	})
}
