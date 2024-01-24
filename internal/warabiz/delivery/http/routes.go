package http

import (
	"github.com/gofiber/fiber/v2"
)

func MapWarabizRoutes(r fiber.Router, handler WarabizHandler) {

	warabiz := r.Group("/warabiz")

	warabiz.Post("", handler.CreateWarabiz)
	warabiz.Get("/id/:id", handler.GetWarabizByID)
	warabiz.Get("/all", handler.GetAllWarabiz)
	warabiz.Put("", handler.UpdateWarabiz)
	warabiz.Delete("/id/:id", handler.DeleteWarabiz)
}
