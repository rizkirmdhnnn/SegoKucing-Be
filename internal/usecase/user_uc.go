package usecase

import (
	"context"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rizkirmdhnnn/segokucing-be/internal/entity"
	"github.com/rizkirmdhnnn/segokucing-be/internal/model"
	"github.com/rizkirmdhnnn/segokucing-be/internal/repository"
	"github.com/rizkirmdhnnn/segokucing-be/internal/utils"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	Validate       *validator.Validate
	UserRepository *repository.UserRepository
	Viper          *viper.Viper
}

func NewUserUseCase(validate *validator.Validate, userRepository *repository.UserRepository, viper *viper.Viper) *UserUseCase {
	return &UserUseCase{
		Validate:       validate,
		UserRepository: userRepository,
		Viper:          viper,
	}
}

func (c *UserUseCase) CreateUser(ctx context.Context, request *model.RegisterUserRequest) (*model.DataUserWithToken, error) {
	// Validate request
	err := c.Validate.Struct(request)
	if err != nil {
		log.Println(err)
		return nil, fiber.ErrBadRequest
	}

	// Check if user already exists
	isRegistered, err := c.UserRepository.IsUserRegistered(request.CredentialValue)
	if err != nil {
		log.Println(err)
		return nil, fiber.ErrInternalServerError
	}
	if isRegistered {
		return nil, fiber.NewError(fiber.StatusConflict, fmt.Sprintf("User with %s already exists", request.CredentialValue))
	}

	// Hash password
	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return nil, fiber.ErrInternalServerError
	}

	// Create user
	user := &entity.Users{
		Name:     request.Name,
		Password: string(password),
	}

	// Set credential
	if request.CredentialType == "email" {
		user.Email = request.CredentialValue
	}
	if request.CredentialType == "phone" {
		user.Phone = request.CredentialValue
	}

	// Save user
	err = c.UserRepository.Create(user)
	if err != nil {
		log.Println(err)
		return nil, fiber.ErrInternalServerError
	}

	// Generate JWT token
	token, err := utils.GenerateToken(int(user.ID), c.Viper.GetInt("JWT_EXPIRATION"), c.Viper.GetString("JWT_SECRET"))
	if err != nil {
		log.Println(err)
		return nil, fiber.ErrInternalServerError
	}

	// Return response
	return &model.DataUserWithToken{
		Phone:       user.Phone,
		Email:       user.Email,
		Name:        user.Name,
		AccessToken: token,
	}, nil

}

func (c *UserUseCase) Login(ctx context.Context, request *model.LoginUserRequest) (*model.DataUserWithToken, error) {
	// Validate request
	err := c.Validate.Struct(request)
	if err != nil {
		log.Println(err)
		return nil, fiber.ErrBadRequest
	}

	// check request credential type
	var user *entity.Users

	if request.CredentialType == "email" {
		user, _ = c.UserRepository.GetUserByEmail(request.CredentialValue)
		if user == nil {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "User not found")
		}
	}

	if request.CredentialType == "phone" {
		user, _ = c.UserRepository.GetUserByPhone(request.CredentialValue)
		if user == nil {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "User not found")
		}
	}

	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		log.Println(err)
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid password")
	}

	// Generate JWT token
	token, err := utils.GenerateToken(int(user.ID), c.Viper.GetInt("JWT_EXPIRATION"), c.Viper.GetString("JWT_SECRET"))
	if err != nil {
		log.Println(err)
		return nil, fiber.ErrInternalServerError
	}

	// Return response
	return &model.DataUserWithToken{
		Phone:       user.Phone,
		Email:       user.Email,
		Name:        user.Name,
		AccessToken: token,
	}, nil
}
