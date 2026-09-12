package cloudinary

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"

	go_cloudinary "github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/ffajarpratama/gommerce-api/config"
	"github.com/ffajarpratama/gommerce-api/constant"
	"github.com/ffajarpratama/gommerce-api/lib/custom_error"
)

type Cloudinary struct {
	cnf      *config.Config
	instance *go_cloudinary.Cloudinary
}

type CloudinaryUploadRes struct {
	PublicID  string `json:"public_id"`
	SecureURL string `json:"secure_url"`
}

func NewClient(cnf *config.Config) (*Cloudinary, error) {
	cld, err := go_cloudinary.NewFromParams(cnf.Cloudinary.CloudName, cnf.Cloudinary.APIKey, cnf.Cloudinary.APISecret)
	if err != nil {
		return nil, err
	}

	log.Println("[cloudinary-connected]")

	return &Cloudinary{
		cnf:      cnf,
		instance: cld,
	}, nil
}

func (c *Cloudinary) UploadImage(ctx context.Context, file multipart.File, loc constant.UploadLocation, assetType string) (*CloudinaryUploadRes, error) {
	opts := uploader.UploadParams{
		ResourceType: assetType,
		Folder:       fmt.Sprintf("/%s/%s", c.cnf.Cloudinary.RootDir, loc),
	}

	res, err := c.instance.Upload.Upload(ctx, file, opts)
	if err != nil {
		return nil, err
	}

	if res.Error.Message != "" {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusBadRequest,
			Message:  fmt.Sprintf("[cloudinary-error] %v", res.Error.Message),
		})

		return nil, err
	}

	_res, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		return nil, err
	}

	fmt.Printf("string(_res): %v\n", string(_res))

	resp := &CloudinaryUploadRes{
		PublicID:  res.PublicID,
		SecureURL: res.SecureURL,
	}

	return resp, nil
}
