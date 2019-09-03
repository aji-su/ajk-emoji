package usecase

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws/awserr"

	"github.com/theboss/ajk-emoji/ajk-func/src/model"
)

type Download struct {
	store
}

func NewDownload(store store) *Show {
	return &Show{
		store: store,
	}
}

func (u *Show) Download(reqID string) (string, error) {
	tarfilekey := reqID + ".tar.gz"
	url := fmt.Sprintf("%s/%s",
		u.store.GetObjectURLPrefix(),
		tarfilekey,
	)
	_, err := u.store.Head(tarfilekey)
	if err != nil {
		if awsErr, ok := err.(awserr.RequestFailure); ok {
			if awsErr.StatusCode() == http.StatusNotFound {
				dir, err := ioutil.TempDir("", "temp")
				if err != nil {
					return "", err
				}
				defer os.RemoveAll(dir)
				fpath := filepath.Join(dir, tarfilekey)
				log.Printf("Creating file: %s", fpath)
				file, err := os.Create(fpath)
				if err != nil {
					return "", err
				}
				defer file.Close()
				if err := u.generateGzip(reqID, file); err != nil {
					return "", err
				}
				if err := u.store.PutFile(fpath, tarfilekey); err != nil {
					return "", err
				}
				return url, nil
			}
		}
		return "", err
	}
	return url, nil
}

func (u *Show) generateGzip(reqID string, dest io.Writer) error {
	jobj, err := u.store.Get(reqID + "/metadata.json")
	if err != nil {
		return err
	}
	defer jobj.Body.Close()
	var meta model.Metadata
	if err := json.NewDecoder(jobj.Body).Decode(&meta); err != nil {
		return err
	}

	zw := gzip.NewWriter(dest)
	defer zw.Close()

	tw := tar.NewWriter(zw)
	defer tw.Close()

	for _, row := range meta.Emojis {
		for _, col := range row {
			obj, err := u.store.Get(col.Key)
			if err != nil {
				log.Printf("error on GetObject: %s; %#v", col.Key, err)
				return err
			}
			defer obj.Body.Close()

			hdr := &tar.Header{
				Name: path.Base(col.Key),
				Mode: 0644,
				Size: int64(obj.ContentLength),
			}
			if err := tw.WriteHeader(hdr); err != nil {
				log.Printf("error on WriteHeader: %#v", err)
				return err
			}
			if _, err := io.Copy(tw, obj.Body); err != nil {
				log.Printf("error on copy: %#v", err)
				return err
			}
		}
	}
	return nil
}
