package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"warabiz/api/config"
	_ "warabiz/api/docs/swagger"
	"warabiz/api/pkg/infra/db"
	"warabiz/api/pkg/infra/logger"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

const (
	certFile        = "ssl/server-cert.pem"
	keyFile         = "ssl/server-key.pem"
	shutdownTimeOut = 5
)

// * Server struct
type Server struct {
	app         *fiber.App
	cfg         *config.Config
	dbList      []db.DatabaseAccount
	redisClient *redis.Client
	mongoDB     *string
	logger      logger.Logger
}

// * NewServer New Server constructor
func NewServer(cfg *config.Config, dbList []db.DatabaseAccount, redisClient *redis.Client, mongoClient *string, logger logger.Logger) *Server {

	//* Initial Engine
	engine := html.New("./pkg/views", ".html")

	//* Initial Fiber App
	app := fiber.New(fiber.Config{
		AppName:               cfg.Server.Name,
		ServerHeader:          "Go Fiber",
		Views:                 engine,
		BodyLimit:             cfg.Server.BodyLimit * 1024 * 1024,
		DisableStartupMessage: false,
	})

	return &Server{app: app, cfg: cfg, dbList: dbList, redisClient: redisClient, mongoDB: mongoClient, logger: logger}
}

func (s *Server) Run() error {
	if s.cfg.Server.SSL {

		if err := s.MapHandlers(s.app); err != nil {
			return err
		}

		go func() {
			s.logger.Infof("Server is listening on PORT: %s", s.cfg.Server.Port)
			if err := s.app.ListenTLS(fmt.Sprintf(":%s", s.cfg.Server.Port), certFile, keyFile); err != nil {
				s.logger.Fatalf("Error starting Server: ", err)
			}
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

		<-quit

		s.logger.Info("Server Exited Properly")
		return s.app.ShutdownWithTimeout(shutdownTimeOut * time.Second)
	}

	if err := s.MapHandlers(s.app); err != nil {
		return err
	}

	go func() {
		s.logger.Infof("Server is listening on PORT: %s", s.cfg.Server.Port)
		if err := s.app.Listen(fmt.Sprintf(":%s", s.cfg.Server.Port)); err != nil {
			s.logger.Fatalf("Error starting Server: ", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	s.logger.Info("Server Exited Properly")
	return s.app.ShutdownWithTimeout(shutdownTimeOut * time.Second)
}
