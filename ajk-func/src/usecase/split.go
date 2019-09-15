package usecase

import (
	"encoding/json"
	"fmt"
	"image"

	"github.com/theboss/ajk-emoji/ajk-func/src/constant"
	"github.com/theboss/ajk-emoji/ajk-func/src/model"
	"golang.org/x/sync/errgroup"
)

// Split represents usecase of image splitting
type Split struct {
	store
	concurrency int
}

// NewSplit returns instance of split usecase
func NewSplit(store store) *Split {
	return &Split{
		store:       store,
		concurrency: constant.Concurrency,
	}
}

// SplitAndPut splits image into number of pieces, and put into s3
func (u *Split) SplitAndPut(reqID string, reqBody *model.RequestBody) error {
	img, err := model.NewImage(
		"original",
		reqID,
		reqBody.ImageAsDataURL,
	)
	if err != nil {
		return err
	}
	if err := u.store.PutImage(img); err != nil {
		return err
	}

	rct := img.Source.Bounds()
	psize := rct.Dx() / reqBody.Xsplit
	ysplit := rct.Dy() / psize

	queue := make(chan *model.Image)
	eg := errgroup.Group{}
	for i := 0; i < u.concurrency; i++ {
		eg.Go(func() error {
			return u.worker(queue)
		})
	}

	var emojis [][]*model.Emoji
	for i := 0; i < ysplit; i++ {
		var es []*model.Emoji
		y := i * psize
		for j := 0; j < reqBody.Xsplit; j++ {
			x := j * psize
			piece := img.NewPiece(
				fmt.Sprintf(reqBody.FnamePrefix+"%02d%02d", i, j),
				image.Rect(x, y, x+psize, y+psize),
			)
			url := fmt.Sprintf("%s/%s",
				u.store.GetObjectURLPrefix(),
				piece.GetFullName(),
			)
			es = append(es, &model.Emoji{
				Shortcode: piece.Name,
				Key:       piece.GetFullName(),
				URL:       url,
			})
			queue <- piece
		}
		emojis = append(emojis, es)
	}
	close(queue)

	if err := eg.Wait(); err != nil {
		return err
	}

	b, err := json.Marshal(&model.Metadata{
		OriginURL: fmt.Sprintf("%s/%s",
			u.store.GetObjectURLPrefix(),
			img.GetFullName()),
		Emojis: emojis,
	})
	if err != nil {
		return err
	}
	return u.store.Put(reqID+"/metadata.json", b)
}

func (u *Split) worker(queue <-chan *model.Image) error {
	for piece := range queue {
		if err := u.store.PutImage(piece); err != nil {
			return err
		}
	}
	return nil
}
