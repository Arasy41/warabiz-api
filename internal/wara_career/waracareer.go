package wara_career

import (
	"github.com/gofiber/fiber/v2"

	"warabiz/api/config"
	"warabiz/api/internal/wara_career/delivery/http"
	"warabiz/api/internal/wara_career/repository"
	"warabiz/api/internal/wara_career/usecase"

	// "warabiz/api/internal/middleware"
	"warabiz/api/pkg/infra/db"
	"warabiz/api/pkg/infra/logger"
)

type WaraCareer struct {
	Repo    repository.Repository
	Usecase usecase.Usecase
	Handler http.WaraCareerHandler
}

func NewWaraCareer(cfg *config.Config, dbList []db.DatabaseAccount, log logger.Logger) WaraCareer {
	repo := repository.NewWaraCareerRepo(dbList, log)
	uc := usecase.NewWaraCareerUsecase(repo, cfg, dbList, log)
	handler := http.NewWaraCareerHandler(uc, cfg, log)
	return WaraCareer{Repo: repo, Usecase: uc, Handler: handler}
}

func NewRoutes(r fiber.Router, handler http.WaraCareerHandler) {
	http.MapWaraCareerRoutes(r, handler)
}
