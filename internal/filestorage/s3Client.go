package filestorage

import (
	"context"
	"dating_service/configs"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
)

func NewS3Client(conf *configs.Config) *s3.Client {
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{URL: conf.S3.Endpoint, SigningRegion: conf.S3.Region}, nil
	})
	awsCfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(customResolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(conf.S3.AccessKeyID, conf.S3.SecretAccessKey, "")),
		config.WithRegion(conf.S3.Region),
	)
	if err != nil {
		log.Fatalf("Ошибка конфигурации AWS SDK: %v", err)
	}
	return s3.NewFromConfig(awsCfg)
}
