package main

import (
	"log"
	"os"
	"strings"

	"warabiz/api/config"
	"warabiz/api/internal/server"
	"warabiz/api/pkg/infra/db"
	"warabiz/api/pkg/infra/logger"
	"warabiz/api/pkg/infra/redis"
)

//? Swagger
// @title Company Profile
// @version 1.0.0
// @description company profile API
// @termsOfService http://swagger.io/terms/

// @contact.name Nsrvel
// @contact.url https://github.com/nsrvel
// @contact.email putra1business@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath
// @schemes http
func main() {

	log.Println("Starting api server")

	//* ====================== Config ======================
	cfg := config.InitConfig(strings.ToLower(os.Getenv("config")))

	//* ====================== Logger ======================
	appLogger := logger.NewApiLogger(cfg)

	appLogger.InitLogger()
	appLogger.Infof("Name: %s, AppVersion: %s, LogLevel: %s, Environtment: %s, SSL: %v", cfg.Server.Name, cfg.Server.AppVersion, cfg.Logger.Level, cfg.Server.Env, cfg.Server.SSL)

	//* ====================== Connection ======================

	//* SQL DB
	var dbList []db.DatabaseAccount

	comproDB := db.NewSqlxDBConnection(&cfg.Connection.Warabiz, appLogger)
	dbList = append(dbList, db.DatabaseAccount{Source: cfg.Connection.Warabiz.DriverSource, SqlDB: comproDB})
	defer comproDB.Close()

	//* Redis
	redisClient := redis.NewRedisClient(&cfg.Connection.Redis)

	defer redisClient.Close()
	appLogger.Info("Redis connected")

	//* Mongo
	// appLogger.Info("MongoDB connected")

	//* ====================== Running Server ======================
	s := server.NewServer(cfg, dbList, redisClient, nil, appLogger)
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
