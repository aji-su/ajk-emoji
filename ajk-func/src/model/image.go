package model

import (
	"bytes"
	"fmt"
	"image"
	"io/ioutil"

	"github.com/disintegration/imaging"
	"github.com/vincent-petithory/dataurl"
)

type Image struct {
	Name      string
	Extension string
	Dirname   string
	Source    image.Image
}

func NewImage(name, dirname, imageAsDataURL string) (*Image, error) {
	dataURL, err := dataurl.DecodeString(imageAsDataURL)
	if err != nil {
		return nil, err
	}

	src, err := imaging.Decode(bytes.NewReader(dataURL.Data))
	if err != nil {
		return nil, err
	}

	return &Image{
		Name:      name,
		Extension: dataURL.MediaType.Subtype,
		Dirname:   dirname,
		Source:    src,
	}, nil
}

func (img *Image) GetFullName() string {
	return fmt.Sprintf("%s/%s.%s", img.Dirname, img.Name, img.Extension)
}

func (img *Image) GetBytes() ([]byte, error) {
	w := bytes.NewBuffer([]byte{})
	if err := imaging.Encode(w, img.Source, imaging.PNG); err != nil {
		return nil, err
	}
	return ioutil.ReadAll(w)
}

func (img *Image) NewPiece(name string, rect image.Rectangle) *Image {
	return &Image{
		Name:      name,
		Extension: img.Extension,
		Dirname:   img.Dirname,
		Source:    imaging.Crop(img.Source, rect),
	}
}
