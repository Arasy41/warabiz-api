package system

import (
	"warabiz/api/config"
	"warabiz/api/internal/system/delivery/http"
	"warabiz/api/pkg/infra/logger"
)

type System struct {
	Handler http.SystemHandler
}

func NewSystem(conf *config.Config, logger logger.Logger) System {
	handler := http.NewSystemHandler(conf, logger)
	return System{Handler: handler}
}
