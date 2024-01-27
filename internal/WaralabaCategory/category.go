package waralabacategory

import (
	"github.com/gofiber/fiber/v2"

	"warabiz/api/config"
	"warabiz/api/internal/WaralabaCategory/delivery/http"
	"warabiz/api/internal/WaralabaCategory/repository"
	"warabiz/api/internal/WaralabaCategory/usecase"

	// "warabiz/api/internal/middleware"
	"warabiz/api/pkg/infra/db"
	"warabiz/api/pkg/infra/logger"
)

type Category struct {
	Repo    repository.Repository
	Usecase usecase.Usecase
	Handler http.CategoryHandler
}

func NewCategory(cfg *config.Config, dbList []db.DatabaseAccount, log logger.Logger) Category {
	repo := repository.NewCategoryRepo(dbList, log)
	uc := usecase.NewCategoryUsecase(repo, cfg, dbList, log)
	handler := http.NewCategoryHandler(uc, cfg, log)
	return Category{Repo: repo, Usecase: uc, Handler: handler}
}

func NewRoutes(r fiber.Router, handler http.CategoryHandler) {
	http.MapCategoryRoutes(r, handler)
}
