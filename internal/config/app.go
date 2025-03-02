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
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Bucket   *minio.Client
	Validate *validator.Validate
	Config   *viper.Viper
	Logger   *logrus.Logger
}

func Bootstrap(config *BootstrapConfig) {
	// Repository
	userRepository := repository.NewUserRepository(config.DB)
	postRepository := repository.NewPostRepository(config.DB)
	tagsRepository := repository.NewTagRepository(config.DB)
	commentRepository := repository.NewCommentRepository(config.DB)
	friendRepository := repository.NewFriendRepository(config.DB)
	fileRepository := repository.NewFileRepository(config.Bucket, config.Config.GetString("S3_BUCKET_NAME"))
	config.Logger.Info("Repository initialized")

	//Usecase
	userUseCase := usecase.NewUserUseCase(userRepository, config.Validate, config.Config, config.Logger)
	postUseCase := usecase.NewPostUseCase(postRepository, tagsRepository, config.Validate, config.Config, config.Logger)
	commentUseCase := usecase.NewCommentUseCase(commentRepository, friendRepository, postRepository, config.Validate, config.Config)
	friendUseCase := usecase.NewFriendUsecase(friendRepository, userRepository, config.Validate, config.Config, config.Logger)
	fileUseCase := usecase.NewFileUsecase(fileRepository)
	config.Logger.Info("Usecase initialized")

	//Controller
	userController := controller.NewUserController(userUseCase, config.Logger)
	PostController := controller.NewPostController(postUseCase, config.Logger)
	commentController := controller.NewCommentController(commentUseCase)
	friendController := controller.NewFriendController(friendUseCase, config.Logger)
	fileController := controller.NewFileController(fileUseCase)
	config.Logger.Info("Controller initialized")

	// // Middleware
	authMiddleware := middleware.NewAuth(config.Config, config.Logger)
	config.Logger.Info("Middleware initialized")

	// // Route
	routeConfig := route.RouteConfig{
		App:               config.App,
		UserController:    userController,
		PostController:    PostController,
		CommentController: commentController,
		FriendController:  friendController,
		FileController:    fileController,
		AuthMiddleware:    authMiddleware,
	}

	routeConfig.Setup()
}
