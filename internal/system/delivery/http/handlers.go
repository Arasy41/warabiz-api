package http

import (
	"net/http"
	"warabiz/api/config"
	"warabiz/api/pkg/http/exception"
	"warabiz/api/pkg/infra/logger"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

type SystemHandler struct {
	cfg    *config.Config
	logger logger.Logger
}

func NewSystemHandler(cfg *config.Config, logger logger.Logger) SystemHandler {
	return SystemHandler{
		cfg:    cfg,
		logger: logger,
	}
}

func (h SystemHandler) Root(c *fiber.Ctx) error {
	return c.Render("root", fiber.Map{"Title": h.cfg.Server.Name})
}

func (h SystemHandler) Swagger(c *fiber.Ctx) error {
	return swagger.HandlerDefault(c)
}

func (h SystemHandler) NotFound(c *fiber.Ctx) error {
	exc := exception.NewException(c, h.logger)
	return exc.WriteErrorResponse(http.StatusNotFound, "Ups! Kami tidak dapat menemukan apa yang Anda cari.", nil)
}
