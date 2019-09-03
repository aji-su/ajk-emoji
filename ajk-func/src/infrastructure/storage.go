package infrastructure

import (
	"bytes"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/theboss/ajk-emoji/ajk-func/src/model"
)

type Storage struct {
	bucketName string
	urlPrefix  string
	client     *s3.S3
}

func NewStorage() *Storage {
	return newStorage(
		os.Getenv("S3_ENDPOINT"),
		os.Getenv("S3_BUCKET_NAME"),
		fmt.Sprintf(
			"https://s3-%s.amazonaws.com/%s",
			os.Getenv("AWS_DEFAULT_REGION"),
			os.Getenv("S3_BUCKET_NAME"),
		),
	)
}

func newStorage(ep, bn, up string) *Storage {
	config := aws.NewConfig().WithS3ForcePathStyle(true)
	if ep != "" {
		config = config.WithEndpoint(ep)
	}
	s := s3.New(session.New(), config)
	return &Storage{
		bucketName: bn,
		urlPrefix:  up,
		client:     s,
	}
}

// TODO: temporary
func (s *Storage) Client() *s3.S3 {
	return s.client
}

func (s *Storage) GetObjectURLPrefix() string {
	return s.urlPrefix
}

func (s *Storage) Put(key string, b []byte) error {
	_, err := s.client.PutObject(&s3.PutObjectInput{
		Body:   bytes.NewReader(b),
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(key),
	})
	return err
}

func (s *Storage) PutImage(img *model.Image) error {
	b, err := img.GetBytes()
	if err != nil {
		return err
	}
	return s.Put(img.GetFullName(), b)
}

func (s *Storage) PutFile(fpath, key string) error {
	f, err := os.Open(fpath)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = s.client.PutObject(&s3.PutObjectInput{
		Body:   f,
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(key),
	})
	return err
}

func (s *Storage) FindByPrefix(prefix string) ([]string, error) {
	output, err := s.client.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(s.bucketName),
		Prefix: aws.String(prefix),
	})
	if err != nil {
		return nil, err
	}
	var keys []string
	for _, item := range output.Contents {
		keys = append(keys, *item.Key)
	}
	return keys, nil
}

func (s *Storage) Get(key string) (*model.StoreObject, error) {
	o, err := s.client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	return &model.StoreObject{
		Body:          o.Body,
		ContentLength: *o.ContentLength,
	}, nil
}

func (s *Storage) Head(key string) (*s3.HeadObjectOutput, error) {
	o, err := s.client.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	return o, nil
}
