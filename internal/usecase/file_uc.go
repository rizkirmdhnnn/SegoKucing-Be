package usecase

import (
	"context"
	"mime/multipart"
	"path/filepath"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rizkirmdhnnn/segokucing-be/internal/model"
	"github.com/rizkirmdhnnn/segokucing-be/internal/repository"
)

type FileUsecase struct {
	fileRepo *repository.FileRepository
}

func NewFileUsecase(fileRepo *repository.FileRepository) *FileUsecase {
	return &FileUsecase{
		fileRepo: fileRepo,
	}
}

func (f *FileUsecase) UploadImageProfile(ctx context.Context, file *multipart.FileHeader) (*model.UploadImageProfileResponse, error) {
	// check if file is nil
	if file == nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "No file uploaded")
	}

	// validate extension
	// only accept jpeg and png
	if file.Header.Get("Content-Type") != "image/jpeg" && file.Header.Get("Content-Type") != "image/png" {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid image type")
	}
	// validate image size
	// > 2MB
	if file.Size > 2*1024*1024 {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Image size too large")
	}

	// < 10KB
	if file.Size < 10*1024 {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Image size too small")
	}

	// get user id from context
	userID := ctx.Value("user_id").(int64)

	// get extension
	extension := filepath.Ext(file.Filename)

	// set new filename
	newFilename := "profile/" + strconv.FormatInt(userID, 10) + extension

	// upload image
	fileURL, err := f.fileRepo.FileUpload(ctx, file, newFilename)
	if err != nil {
		return nil, err
	}

	res := &model.UploadImageProfileResponse{
		ImageUrl: fileURL,
	}

	return res, nil
}
