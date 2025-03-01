package config

import (
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewBucket(viper *viper.Viper, log *logrus.Logger) *minio.Client {
	endpoint := viper.GetString("S3_ENDPOINT")
	accessKeyID := viper.GetString("S3_ID")
	secretAccessKey := viper.GetString("S3_SECRET_KEY")
	bucketName := viper.GetString("S3_BUCKET_NAME")
	useSSL := viper.GetBool("S3_USE_SSL")

	// Log bucket configuration
	log.Infof("S3 Endpoint: %s", endpoint)
	log.Infof("S3 Access Key ID: %s", accessKeyID)
	log.Infof("S3 Bucket Name: %s", bucketName)
	log.Infof("S3 Use SSL: %t", useSSL)

	// Initialize MinIO client
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalf("Error initializing MinIO client: %v", err)
	}

	// Check if bucket exists
	log.Infof("Checking if bucket %s exists", bucketName)
	ctx := context.Background()
	exists, err := minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		log.Fatalf("Error checking if bucket %s exists: %v", bucketName, err)
	}
	if !exists {
		log.Fatalf("Bucket %s does not exist", bucketName)
	}

	return minioClient
}
