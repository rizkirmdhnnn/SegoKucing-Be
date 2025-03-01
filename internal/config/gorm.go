package config

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(cfg *viper.Viper, log *logrus.Logger) *gorm.DB {
	// Get database configuration
	host := cfg.GetString("DB_HOST")
	port := cfg.GetInt("DB_PORT")
	username := cfg.GetString("DB_USERNAME")
	password := cfg.GetString("DB_PASSWORD")
	params := cfg.GetString("DB_PARAMS")
	database := cfg.GetString("DB_NAME")
	idleConnection := viper.GetInt("DB_MAX_IDLE_CONNECTION")
	maxConnection := viper.GetInt("DB_MAX_CONNECTION")
	maxLifeTimeConnection := viper.GetInt("DB_MAX_LIFETIME_CONNECTION")

	// Connect to database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d %s", host, username, password, database, port, params)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(&logrusWriter{Logger: log}, logger.Config{
			SlowThreshold:             time.Second * 5,
			Colorful:                  false,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			LogLevel:                  logger.Info,
		}),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Check connection
	connection, err := db.DB()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Set connection configuration
	connection.SetMaxIdleConns(idleConnection)
	connection.SetMaxOpenConns(maxConnection)
	connection.SetConnMaxLifetime(time.Second * time.Duration(maxLifeTimeConnection))

	return db
}

type logrusWriter struct {
	Logger *logrus.Logger
}

func (l *logrusWriter) Printf(message string, args ...interface{}) {
	l.Logger.Tracef(message, args...)
}
