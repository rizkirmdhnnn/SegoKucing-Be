package main

import (
	"fmt"

	"github.com/rizkirmdhnnn/segokucing-be/internal/config"
)

func main() {
	// Initialize the configuration
	cfg := config.NewViper()

	// Initialize the logger
	log := config.NewLogger(cfg)

	// Initialize the database
	db := config.NewDatabase(cfg, log)

	// Initialize the validator
	validate := config.NewValidator(cfg)

	// Initialize the Minio client
	bucket := config.NewBucket(cfg, log)

	// Initialize the Fiber app
	app := config.NewFiber(cfg)

	// Bootstrap the application
	config.Bootstrap(&config.BootstrapConfig{
		DB:       db,
		App:      app,
		Bucket:   bucket,
		Validate: validate,
		Config:   cfg,
		Logger:   log,
	})

	// Start the server
	webPort := cfg.GetInt("APP_PORT")
	if err := app.Listen(fmt.Sprintf(":%d", webPort)); err != nil {
		panic(err)
	}
}
