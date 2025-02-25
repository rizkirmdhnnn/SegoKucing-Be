package repository

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
)

type FileRepository struct {
	client     *minio.Client
	bucketName string
}

func NewFileRepository(minio *minio.Client, bucketName string) *FileRepository {
	return &FileRepository{
		client:     minio,
		bucketName: bucketName,
	}
}

func (f *FileRepository) FileUpload(ctx context.Context, file *multipart.FileHeader, fileName string) (string, error) {
	// Buka file yang diupload
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Upload file ke MinIO
	_, err = f.client.PutObject(
		ctx,
		f.bucketName,
		fileName,
		src,
		file.Size,
		minio.PutObjectOptions{
			ContentType: file.Header.Get("Content-Type"),
			UserMetadata: map[string]string{
				"x-amz-acl": "public-read",
			},
		},
	)
	if err != nil {
		return "", err
	}

	// Generate URL untuk mengakses file
	fileURL := fmt.Sprintf("%s/%s/%s", f.client.EndpointURL(), f.bucketName, fileName)

	return fileURL, nil
}
