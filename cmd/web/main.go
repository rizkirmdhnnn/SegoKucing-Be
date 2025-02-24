package main

import (
	"fmt"

	"github.com/rizkirmdhnnn/segokucing-be/internal/config"
)

func main() {
	// Initialize the configuration
	cfg := config.NewViper()

	// Initialize the database
	db := config.NewDatabase(cfg)

	// Initialize the validator
	validate := config.NewValidator(cfg)

	// Initialize the Minio client
	bucket := config.NewBucket(cfg)

	// Initialize the Fiber app
	app := config.NewFiber(cfg)

	// Bootstrap the application
	config.Bootstrap(&config.BootstrapConfig{
		DB:       db,
		App:      app,
		Bucket:   bucket,
		Validate: validate,
		Config:   cfg,
	})

	// Start the server
	webPort := cfg.GetInt("APP_PORt")
	if err := app.Listen(fmt.Sprintf(":%d", webPort)); err != nil {
		panic(err)
	}
}
