package getter

import (
	"fmt"
	cg "warabiz/api/pkg/constants/general"

	"github.com/gofiber/fiber/v2"
)

func GetRequestID(c *fiber.Ctx) string {
	return fmt.Sprintf("%v", c.Locals(cg.CtxRequestID))
}
