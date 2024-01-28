package http

import (
	// "warabiz/api/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func MapWaraCareerRoutes(r fiber.Router, handler WaraCareerHandler) {

	cmsWaraCareer := r.Group("/cms/wara-career")

	cmsWaraCareer.Post("", handler.CreateWaraCareer)
	cmsWaraCareer.Get("/id/:id", handler.GetWaraCareerByID)
	cmsWaraCareer.Get("/all", handler.GetAllWaraCareer)
	cmsWaraCareer.Put("", handler.UpdateWaraCareer)
	cmsWaraCareer.Delete("/id/:id", handler.DeleteWaraCareer)
	// cmsWaraCareer.Get("/detail/:id", handler.GetWaraCareerDetails)
}
