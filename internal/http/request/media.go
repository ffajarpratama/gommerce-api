package request

import (
	"mime/multipart"

	"github.com/ffajarpratama/gommerce-api/constant"
)

type CreateMedia struct {
	File      multipart.File
	Header    *multipart.FileHeader
	Location  constant.UploadLocation
	Filename  string
	Mimetype  string
	AssetType string
}
