package config

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)

func NewBucket(viper *viper.Viper) *minio.Client {
	endpoint := viper.GetString("S3_ENDPOINT")
	accessKeyID := viper.GetString("S3_ID")
	secretAccessKey := viper.GetString("S3_SECRET_KEY")
	bucket_name := viper.GetString("S3_BUCKET_NAME")
	useSSL := viper.GetBool("S3_USE_SSL")

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()

	err = minioClient.MakeBucket(ctx, bucket_name, minio.MakeBucketOptions{})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucket_name)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucket_name)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucket_name)
	}

	return minioClient
}
