package usecase

import (
	"context"
	"path/filepath"

	"github.com/ffajarpratama/gommerce-api/internal/http/request"
	"github.com/ffajarpratama/gommerce-api/internal/model"
)

// CreateMedia implements IFaceUsecase.
func (u *Usecase) CreateMedia(ctx context.Context, req *request.CreateMedia) (*model.Media, error) {
	res, err := u.cld.UploadImage(ctx, req.File, req.Location, req.AssetType)
	if err != nil {
		return nil, err
	}

	media := &model.Media{
		Name:     req.Filename,
		Path:     res.PublicID + filepath.Ext(req.Header.Filename),
		Size:     int(req.Header.Size),
		Mimetype: req.Mimetype,
		MediaURL: res.SecureURL,
	}

	err = u.repo.CreateMedia(ctx, media, u.db)
	if err != nil {
		return nil, err
	}

	return media, nil
}
