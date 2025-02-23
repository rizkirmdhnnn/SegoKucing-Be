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
	UserRepository *repository.UserRepository
	Validate       *validator.Validate
	Viper          *viper.Viper
}

func NewUserUseCase(userRepository *repository.UserRepository, validate *validator.Validate, viper *viper.Viper) *UserUseCase {
	return &UserUseCase{
		UserRepository: userRepository,
		Validate:       validate,
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
	token, err := utils.GenerateToken(user.ID, c.Viper)
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
	token, err := utils.GenerateToken(user.ID, c.Viper)
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

// link email
func (c *UserUseCase) LinkEmail(ctx context.Context, request *model.LinkEmailRequest) error {
	userId := ctx.Value("user_id").(int64)

	// Validate request
	err := c.Validate.Struct(request)
	if err != nil {
		log.Println(err)
		return fiber.ErrBadRequest
	}

	// Checi if user already linked email
	user, err := c.UserRepository.GetUserById(int(userId))
	if err != nil {
		log.Println(err)
		return fiber.ErrInternalServerError
	}

	if user.Email != "" {
		return fiber.NewError(fiber.StatusConflict, "User already linked email")
	}

	// Check if user already exists
	isRegistered, err := c.UserRepository.IsUserRegistered(request.Email)
	if err != nil {
		log.Println(err)
		return fiber.ErrInternalServerError
	}
	if isRegistered {
		return fiber.NewError(fiber.StatusConflict, fmt.Sprintf("User with %s already exists", request.Email))
	}

	// Update user
	user.Email = request.Email
	err = c.UserRepository.Update(user)
	if err != nil {
		log.Println(err)
		return fiber.ErrInternalServerError
	}

	return nil
}

func (c *UserUseCase) LinkPhoneNumber(ctx context.Context, request *model.LinkPhoneRequest) error {
	userId := ctx.Value("user_id").(int64)

	// Validate request
	err := c.Validate.Struct(request)
	if err != nil {
		log.Println(err)
		return fiber.ErrBadRequest
	}

	// Checi if user already linked email
	user, err := c.UserRepository.GetUserById(int(userId))
	if err != nil {
		log.Println(err)
		return fiber.ErrInternalServerError
	}

	if user.Phone != "" {
		return fiber.NewError(fiber.StatusConflict, "User already linked phone number")
	}

	// Check if user already exists
	isRegistered, err := c.UserRepository.IsUserRegistered(request.Phone)
	if err != nil {
		log.Println(err)
		return fiber.ErrInternalServerError
	}
	if isRegistered {
		return fiber.NewError(fiber.StatusConflict, fmt.Sprintf("User with %s already exists", request.Phone))
	}

	// Update user
	user.Phone = request.Phone
	err = c.UserRepository.Update(user)
	if err != nil {
		log.Println(err)
		return fiber.ErrInternalServerError
	}

	return nil
}
