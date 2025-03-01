package config

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)

func NewBucket(viper *viper.Viper) *minio.Client {
	// Get values from the config file
	endpoint := viper.GetString("S3_ENDPOINT")
	accessKeyID := viper.GetString("S3_ID")
	secretAccessKey := viper.GetString("S3_SECRET_KEY")
	bucketName := viper.GetString("S3_BUCKET_NAME")
	useSSL := viper.GetBool("S3_USE_SSL")

	// Log the values
	log.Printf("Endpoint: %s\n", endpoint)
	log.Printf("Access Key ID: %s\n", accessKeyID)
	log.Printf("Secret Access Key: %s\n", secretAccessKey)
	log.Printf("Bucket Name: %s\n", bucketName)
	log.Printf("Use SSL: %v\n", useSSL)

	// Initialize MinIO client
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalf("Error initializing MinIO client: %v", err)
	}

	ctx := context.Background()
	// Check if bucket already exists
	exists, err := minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		log.Fatalf("Error checking if bucket exists: %v", err)
	}
	if !exists {
		log.Fatalf("Bucket %s does not exist", bucketName)
	}

	return minioClient
}
