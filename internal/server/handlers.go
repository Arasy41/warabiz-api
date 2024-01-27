package server

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	waralabacategory "warabiz/api/internal/WaralabaCategory"
	"warabiz/api/internal/middleware"
	"warabiz/api/internal/system"
	"warabiz/api/internal/warabiz"
)

func (s *Server) MapHandlers(app *fiber.App) error {

	//* Initial Repoconf
	System := system.NewSystem(s.cfg, s.logger)
	Warabiz := warabiz.NewWarabiz(s.cfg, s.dbList, s.logger)
	Category := waralabacategory.NewCategory(s.cfg, s.dbList, s.logger)

	//* Set middleware constructor
	mw := middleware.NewMiddlewareManager(s.cfg, s.dbList, s.logger)

	//* Root
	app.Get("/", System.Handler.Root)

	//* Swagger
	app.Get("/swagger/*", System.Handler.Swagger)

	//* General Middlewarex
	if strings.ToLower(s.cfg.Server.Env) == "production" {
		app.Use(mw.Recover())
	}
	app.Use(mw.CORS())
	app.Use(mw.RequestID())
	// app.Use(mw.Limitter())
	// app.Use(mw.RequestLog())

	//* Map routes
	warabiz.NewRoutes(app, Warabiz.Handler)
	waralabacategory.NewRoutes(app, Category.Handler)

	//* Not found
	app.All("*", System.Handler.NotFound)

	return nil
}
