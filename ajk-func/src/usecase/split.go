package usecase

import (
	"encoding/json"
	"fmt"
	"image"
	"log"

	"github.com/theboss/ajk-emoji/ajk-func/src/model"
)

type Split struct {
	store
}

func NewSplit(store store) *Split {
	return &Split{
		store: store,
	}
}

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
			if err := u.store.PutImage(piece); err != nil {
				return err
			}
			url := fmt.Sprintf("%s/%s",
				u.store.GetObjectURLPrefix(),
				piece.GetFullName(),
			)
			es = append(es, &model.Emoji{
				Shortcode: piece.Name,
				Key:       piece.GetFullName(),
				URL:       url,
			})
		}
		emojis = append(emojis, es)
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
