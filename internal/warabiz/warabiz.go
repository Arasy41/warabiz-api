package warabiz

import (
	"github.com/gofiber/fiber/v2"

	"warabiz/api/config"
	"warabiz/api/internal/warabiz/delivery/http"
	"warabiz/api/internal/warabiz/repository"
	"warabiz/api/internal/warabiz/usecase"

	// "warabiz/api/internal/middleware"
	"warabiz/api/pkg/infra/db"
	"warabiz/api/pkg/infra/logger"
)

type Warabiz struct {
	Repo    repository.Repository
	Usecase usecase.Usecase
	Handler http.WarabizHandler
}

func NewWarabiz(cfg *config.Config, dbList []db.DatabaseAccount, log logger.Logger) Warabiz {
	repo := repository.NewWarabizRepo(dbList, log)
	uc := usecase.NewWarabizUsecase(repo, cfg, dbList, log)
	handler := http.NewWarabizHandler(uc, cfg, log)
	return Warabiz{Repo: repo, Usecase: uc, Handler: handler}
}

func NewRoutes(r fiber.Router, handler http.WarabizHandler) {
	http.MapWarabizRoutes(r, handler)
}
