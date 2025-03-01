package usecase

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rizkirmdhnnn/segokucing-be/internal/entity"
	"github.com/rizkirmdhnnn/segokucing-be/internal/model"
	"github.com/rizkirmdhnnn/segokucing-be/internal/repository"
	"github.com/rizkirmdhnnn/segokucing-be/internal/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	UserRepository *repository.UserRepository
	Validate       *validator.Validate
	Viper          *viper.Viper
	Log            *logrus.Logger
}

func NewUserUseCase(userRepository *repository.UserRepository, validate *validator.Validate, viper *viper.Viper, log *logrus.Logger) *UserUseCase {
	return &UserUseCase{
		UserRepository: userRepository,
		Validate:       validate,
		Viper:          viper,
		Log:            log,
	}
}

func (c *UserUseCase) CreateUser(ctx context.Context, request *model.RegisterUserRequest) (*model.DataUserWithToken, error) {
	c.Log.Info("Creating user")

	// Validate request
	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Error("Validation failed: ", err)
		return nil, fiber.ErrBadRequest
	}

	// Check if user already exists
	isRegistered, err := c.UserRepository.IsUserRegistered(request.CredentialValue)
	if err != nil {
		c.Log.Error("Error checking user registration: ", err)
		return nil, fiber.ErrInternalServerError
	}
	if isRegistered {
		return nil, fiber.NewError(fiber.StatusConflict, fmt.Sprintf("User with %s already exists", request.CredentialValue))
	}

	// Hash password
	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Log.Error("Error hashing password: ", err)
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
		c.Log.Error("Error saving user: ", err)
		return nil, fiber.ErrInternalServerError
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.ID, c.Viper)
	if err != nil {
		c.Log.Error("Error generating token: ", err)
		return nil, fiber.ErrInternalServerError
	}

	c.Log.Info("User created successfully")

	// Return response
	return &model.DataUserWithToken{
		Phone:       user.Phone,
		Email:       user.Email,
		Name:        user.Name,
		AccessToken: token,
	}, nil
}

func (c *UserUseCase) Login(ctx context.Context, request *model.LoginUserRequest) (*model.DataUserWithToken, error) {
	c.Log.Info("User login attempt")

	// Validate request
	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Error("Validation failed: ", err)
		return nil, fiber.ErrBadRequest
	}

	// Check user credential
	var user *entity.Users

	if request.CredentialType == "email" {
		user, err = c.UserRepository.GetUserByEmail(request.CredentialValue)
	} else {
		user, err = c.UserRepository.GetUserByPhone(request.CredentialValue)
	}

	if err != nil {
		c.Log.Error("Error fetching user: ", err)
		return nil, fiber.ErrInternalServerError
	}

	if user == nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "User not found")
	}

	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		c.Log.Warn("Invalid password for user: ", request.CredentialValue)
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid password")
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.ID, c.Viper)
	if err != nil {
		c.Log.Error("Error generating token: ", err)
		return nil, fiber.ErrInternalServerError
	}

	c.Log.Info("User login successful")

	// Return response
	return &model.DataUserWithToken{
		Phone:       user.Phone,
		Email:       user.Email,
		Name:        user.Name,
		AccessToken: token,
	}, nil
}

func (c *UserUseCase) LinkEmail(ctx context.Context, request *model.LinkEmailRequest) error {
	userId := ctx.Value("user_id").(int64)
	c.Log.WithFields(logrus.Fields{
		"user_id": userId,
	}).Info("Linking email")

	// Validate request
	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Error("Validation failed: ", err)
		return fiber.ErrBadRequest
	}

	// Check if user already linked email
	user, err := c.UserRepository.GetUserById(int(userId))
	if err != nil {
		c.Log.Error("Error fetching user: ", err)
		return fiber.ErrInternalServerError
	}

	if user.Email != "" {
		return fiber.NewError(fiber.StatusConflict, "User already linked email")
	}

	// Check if email is already registered
	isRegistered, err := c.UserRepository.IsUserRegistered(request.Email)
	if err != nil {
		c.Log.Error("Error checking email registration: ", err)
		return fiber.ErrInternalServerError
	}
	if isRegistered {
		return fiber.NewError(fiber.StatusConflict, fmt.Sprintf("User with %s already exists", request.Email))
	}

	// Update user
	user.Email = request.Email
	err = c.UserRepository.Update(user)
	if err != nil {
		c.Log.Error("Error updating user: ", err)
		return fiber.ErrInternalServerError
	}

	c.Log.Info("Email linked successfully")
	return nil
}

func (c *UserUseCase) LinkPhoneNumber(ctx context.Context, request *model.LinkPhoneRequest) error {
	userId := ctx.Value("user_id").(int64)
	c.Log.WithFields(logrus.Fields{
		"user_id": userId,
	}).Info("Linking phone number")

	// Validate request
	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Error("Validation failed: ", err)
		return fiber.ErrBadRequest
	}

	// Check if user already linked email
	user, err := c.UserRepository.GetUserById(int(userId))
	if err != nil {
		c.Log.Error("Error fetching user: ", err)
		return fiber.ErrInternalServerError
	}

	if user.Phone != "" {
		return fiber.NewError(fiber.StatusConflict, "User already linked phone number")
	}

	// Check if email is already registered
	isRegistered, err := c.UserRepository.IsUserRegistered(request.Phone)
	if err != nil {
		c.Log.Error("Error checking email registration: ", err)
		return fiber.ErrInternalServerError
	}
	if isRegistered {
		return fiber.NewError(fiber.StatusConflict, fmt.Sprintf("User with %s already exists", request.Phone))
	}

	// Update user
	user.Phone = request.Phone
	err = c.UserRepository.Update(user)
	if err != nil {
		c.Log.Error("Error updating user: ", err)
		return fiber.ErrInternalServerError
	}

	c.Log.Info("Phone number linked successfully")
	return nil
}
