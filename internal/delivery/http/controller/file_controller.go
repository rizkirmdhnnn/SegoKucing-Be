package controller

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/rizkirmdhnnn/segokucing-be/internal/usecase"
)

type FileController struct {
	fileUsecase *usecase.FileUsecase
}

func NewFileController(fileUsecase *usecase.FileUsecase) *FileController {
	return &FileController{
		fileUsecase: fileUsecase,
	}
}

func (f *FileController) UploadImageProfile(ctx *fiber.Ctx) error {
	// Get User ID from context
	userID, ok := ctx.Locals("user_id").(int64)
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// Get file from form
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	// Upload image
	newCtx := context.WithValue(ctx.UserContext(), "user_id", userID)
	response, err := f.fileUsecase.UploadImageProfile(newCtx, file)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
		return err
	}

	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Image uploaded successfully",
		"data":    response,
	})

	return nil
}
