package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"github.com/rizkirmdhnnn/segokucing-be/internal/delivery/http/controller"
	"github.com/rizkirmdhnnn/segokucing-be/internal/delivery/http/middleware"
	"github.com/rizkirmdhnnn/segokucing-be/internal/delivery/http/route"
	"github.com/rizkirmdhnnn/segokucing-be/internal/repository"
	"github.com/rizkirmdhnnn/segokucing-be/internal/usecase"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Bucket   *minio.Client
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	// Repository
	userRepository := repository.NewUserRepository(config.DB)
	postRepository := repository.NewPostRepository(config.DB)
	tagsRepository := repository.NewTagRepository(config.DB)
	commentRepository := repository.NewCommentRepository(config.DB)
	friendRepository := repository.NewFriendRepository(config.DB)

	//Usecase
	userUseCase := usecase.NewUserUseCase(userRepository, config.Validate, config.Config)
	postUseCase := usecase.NewPostUseCase(postRepository, tagsRepository, config.Validate, config.Config)
	commentUseCase := usecase.NewCommentUseCase(commentRepository, friendRepository, postRepository, config.Validate, config.Config)
	friendUseCase := usecase.NewFriendUsecase(friendRepository, userRepository, config.Validate, config.Config)

	//Controller
	userController := controller.NewUserController(userUseCase)
	PostController := controller.NewPostController(postUseCase)
	commentController := controller.NewCommentController(commentUseCase)
	friendController := controller.NewFriendController(friendUseCase)

	// // Middleware
	authMiddleware := middleware.NewAuth(config.Config)

	// // Route
	routeConfig := route.RouteConfig{
		App:               config.App,
		UserController:    userController,
		PostController:    PostController,
		CommentController: commentController,
		FriendController:  friendController,
		AuthMiddleware:    authMiddleware,
	}

	routeConfig.Setup()
}
