package db

import (
	"strings"
	"time"

	"warabiz/api/config"
	"warabiz/api/pkg/infra/logger"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	gormlog "gorm.io/gorm/logger"
)

func NewSqlxDBConnection(cfg *config.DatabaseAccount, log logger.Logger) *sqlx.DB {

	//* Get DBName from DB source
	dbName := GetDBNameFromDriverSource(cfg.DriverSource)

	db, err := sqlx.Connect(cfg.ServerType, cfg.DriverSource)
	if err != nil {
		log.Fatal("Failed to connect database " + dbName + ", err: " + err.Error())
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime * time.Minute)
	db.SetConnMaxIdleTime(cfg.ConnMaxIdleTime * time.Minute)
	if err = db.Ping(); err != nil {
		log.Fatal("Failed to connect database " + dbName + ", err: " + err.Error())
	}

	log.Info("Connection opened to database " + dbName)
	return db
}

func NewGormDBConnection(cfg *config.DatabaseAccount, log logger.Logger) *gorm.DB {

	//* Get DBName from DB source
	dbName := GetDBNameFromDriverSource(cfg.DriverSource)

	gormConf := &gorm.Config{
		Logger: gormlog.Default.LogMode(gormlog.LogLevel(gormlog.Error)),
		// Logger: logger.Default.LogMode(logger.Silent),
		// SkipDefaultTransaction: true,
	}

	serverType := strings.ToLower(cfg.ServerType)
	var db *gorm.DB
	var err error

	switch serverType {
	case "mysql":
		db, err = gorm.Open(mysql.Open(cfg.DriverSource), gormConf)
	case "postgres":
		db, err = gorm.Open(postgres.Open(cfg.DriverSource), gormConf)
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(cfg.DriverSource), gormConf)
	case "sqlserver":
		db, err = gorm.Open(sqlserver.Open(cfg.DriverSource), gormConf)
	default:
		log.Fatal("failed to connect database " + dbName + ", err: unknown driver")
	}

	if err != nil {
		log.Fatal("failed to connect database " + dbName + ", err: " + err.Error())
	}

	log.Info("connection opened to database " + dbName)
	return db
}
