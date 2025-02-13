package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rizkirmdhnnn/segokucing-be/internal/delivery/http/controller"
	"github.com/rizkirmdhnnn/segokucing-be/internal/delivery/http/route"
	"github.com/rizkirmdhnnn/segokucing-be/internal/repository"
	"github.com/rizkirmdhnnn/segokucing-be/internal/usecase"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	// Repository
	userRepository := repository.NewUserRepository(config.DB)

	// // Usecase
	userUseCase := usecase.NewUserUseCase(config.Validate, userRepository, config.Config)

	// // Controller
	userController := controller.NewUserController(userUseCase)

	// // Middleware
	// authMiddleware := middleware.NewAuth(userUseCase)

	// // Route
	routeConfig := route.RouteConfig{
		App:            config.App,
		UserController: userController,
		// AuthMiddleware: authMiddleware,
	}

	routeConfig.Setup()
}
