package usecase

import (
	"io"
	"log"
)

type Show struct {
	store
}

func NewShow(store store) *Show {
	return &Show{
		store: store,
	}
}

func (u *Show) GetEmojis(reqID string) (io.ReadCloser, error) {
	log.Printf("ReqID=%s", reqID)

	obj, err := u.store.Get(reqID + "/metadata.json")
	if err != nil {
		return nil, err
	}
	return obj.Body, nil
}
