package constant

type UploadLocation string

const (
	UploadLocationAvatar  UploadLocation = "avatar"
	UploadLocationProduct UploadLocation = "product"
)

var AllowedUploadLocation = map[UploadLocation]bool{
	UploadLocationAvatar:  true,
	UploadLocationProduct: true,
}

var MimetypeWhitelist = map[string]bool{
	"image/jpg":  true,
	"image/jpeg": true,
	"image/png":  true,
	"image/webp": true,
}

func GetAssetType(mimetype string) string {
	switch mimetype {
	case "image/jpg", "image/jpeg", "image/png", "image/webp":
		return "image"
	default:
		return "auto"
	}
}
