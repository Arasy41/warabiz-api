package middleware

import (
	"warabiz/api/config"
	"warabiz/api/pkg/infra/db"
	"warabiz/api/pkg/infra/logger"
)

type MiddlewareManager struct {	
	cfg           *config.Config
	dbList        []db.DatabaseAccount
	logger        logger.Logger
}

func NewMiddlewareManager(cfg *config.Config, dbList []db.DatabaseAccount, logger logger.Logger) *MiddlewareManager {
	return &MiddlewareManager{		
		cfg:           cfg,
		dbList:        dbList,
		logger:        logger,
	}
}
