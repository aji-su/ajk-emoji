package model

import "io"

type StoreObject struct {
	Body          io.ReadCloser
	ContentLength int64
}
