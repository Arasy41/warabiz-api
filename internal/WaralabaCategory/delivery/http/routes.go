package http

import (
	// "warabiz/api/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func MapCategoryRoutes(r fiber.Router, handler CategoryHandler) {

	cmsCategories := r.Group("/cms/category")

	cmsCategories.Post("", handler.CreateCategory)
	cmsCategories.Get("/id/:id", handler.GetCategoryByID)
	cmsCategories.Get("/all", handler.GetAllCategory)
	cmsCategories.Put("", handler.UpdateCategory)
	cmsCategories.Delete("/id/:id", handler.DeleteCategory)
	// cmsCategories.Get("/detail/:id", handler.GetCategoryDetails)
}
