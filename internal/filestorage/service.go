package filestorage

import (
	"context"
	"dating_service/configs"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"io"
	"path/filepath"
)

type S3FileStorage struct {
	client     *s3.Client
	bucketName string
	endpoint   string
}

func NewS3FileStorage(client *s3.Client, conf *configs.Config) *S3FileStorage {
	return &S3FileStorage{
		endpoint:   conf.S3.Endpoint,
		bucketName: conf.S3.BucketName,
		client:     client,
	}
}

func (s *S3FileStorage) UploadFile(ctx context.Context, file io.Reader, fileName string) (string, string, error) {
	objectKey := uuid.NewString() + filepath.Ext(fileName)
	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(objectKey),
		Body:   file,
	})
	if err != nil {
		return "", "", ErrUploadFile
	}
	fileURL := fmt.Sprintf("%s/%s/%s", s.endpoint, s.bucketName, objectKey)
	return fileURL, objectKey, nil
}

func (s *S3FileStorage) Delete(ctx context.Context, objectKey string) error {
	_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return ErrDeleteFile
	}
	return nil
}
