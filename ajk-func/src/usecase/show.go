package usecase

import (
	"io"
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
	obj, err := u.store.Get(reqID + "/metadata.json")
	if err != nil {
		return nil, err
	}
	return obj.Body, nil
}
