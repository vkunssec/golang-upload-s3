package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func main() {
	ctx := context.Background()

	var bucket, key string

	flag.StringVar(&bucket, "b", "", "Bucket name.")
	flag.StringVar(&key, "k", "", "Object key name.")
	flag.Parse()

	log.Println(bucket, key)

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	cfg.Region = "sa-east-1"

	s3Client := s3.NewFromConfig(cfg)

	file, err := os.Open("file_test.txt")
	if err != nil {
		log.Fatal("Failed opening file", "file_test.txt", err)
	}

	input := &s3.PutObjectInput{
		Key:    aws.String(key),
		Bucket: aws.String(bucket),
		Body:   file,
		ACL:    types.ObjectCannedACLPublicRead,
	}

	if _, err := s3Client.PutObject(ctx, input); err != nil {
		log.Fatal("fn UploadFile %w", err)
	}

	log.Printf("successfully uploaded file to %s/%s\n", bucket, key)
	log.Println()
	file.Close()
}
