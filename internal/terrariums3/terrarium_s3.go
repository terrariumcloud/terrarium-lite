package terrariums3

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type TerrariumS3Storage struct {
	Region  string
	Service *s3.Client
	config  aws.Config
}

func (s *TerrariumS3Storage) Init() error {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(s.Region))
	if err != nil {
		return err
	}
	s.Service = s3.NewFromConfig(cfg)
	s.config = cfg
	return nil
}

func (s *TerrariumS3Storage) FetchModuleSource(ctx context.Context, bucket string, key string) error {
	data, err := s.Service.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return err
	}
	log.Println(data.ContentType)
	log.Println(data.ContentDisposition)
	return nil
}

func New(region string) (*TerrariumS3Storage, error) {
	s := &TerrariumS3Storage{
		Region: region,
	}
	err := s.Init()
	if err != nil {
		return nil, err
	}
	return s, nil
}
