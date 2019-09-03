package usecase

import (
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/theboss/ajk-emoji/ajk-func/src/model"
)

type store interface {
	Put(string, []byte) error
	PutImage(*model.Image) error
	GetObjectURLPrefix() string
	FindByPrefix(string) ([]string, error)
	Get(string) (*model.StoreObject, error)
	Head(string) (*s3.HeadObjectOutput, error)
	PutFile(string, string) error
}
