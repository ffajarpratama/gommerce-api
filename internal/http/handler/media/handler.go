package media

import (
	"errors"
	"net/http"

	"github.com/ffajarpratama/gommerce-api/config"
	"github.com/ffajarpratama/gommerce-api/constant"
	"github.com/ffajarpratama/gommerce-api/internal/http/middleware"
	"github.com/ffajarpratama/gommerce-api/internal/http/request"
	"github.com/ffajarpratama/gommerce-api/internal/http/response"
	"github.com/ffajarpratama/gommerce-api/internal/usecase"
	"github.com/ffajarpratama/gommerce-api/lib/custom_error"
	"github.com/go-chi/chi/v5"
)

func NewHTTPHandler(cnf *config.Config, uc usecase.IFaceUsecase) http.Handler {
	r := chi.NewRouter()

	r.Group(func(private chi.Router) {
		private.Use(middleware.Authorize(cnf.JWT.Secret))
		private.Post("/", func(w http.ResponseWriter, r *http.Request) {
			r.Body = http.MaxBytesReader(w, r.Body, constant.FileUploadMaxSize)
			if err := r.ParseMultipartForm(constant.FileUploadMaxSize); err != nil {
				err = custom_error.SetCustomError(&custom_error.ErrorContext{
					HTTPCode: http.StatusBadRequest,
					Message:  "file size exceeds 2MB",
				})

				response.Error(w, err)
				return
			}

			file, header, err := r.FormFile("file")
			if err != nil {
				if errors.Is(err, http.ErrMissingFile) {
					err = custom_error.SetCustomError(&custom_error.ErrorContext{
						HTTPCode: http.StatusBadRequest,
						Message:  "file cannot be empty",
					})

					response.Error(w, err)
					return
				}

				response.Error(w, err)
				return
			}

			location := constant.UploadLocation(r.FormValue("location"))
			mimetype := header.Header.Get("Content-Type")

			if !constant.MimetypeWhitelist[mimetype] {
				err = custom_error.SetCustomError(&custom_error.ErrorContext{
					HTTPCode: http.StatusBadRequest,
					Message:  "file extension not allowed",
				})

				response.Error(w, err)
				return
			}

			if !constant.AllowedUploadLocation[location] {
				err = custom_error.SetCustomError(&custom_error.ErrorContext{
					HTTPCode: http.StatusBadRequest,
					Message:  "upload location must be one of [avatar, product]",
				})

				response.Error(w, err)
				return
			}

			req := &request.CreateMedia{
				File:      file,
				Header:    header,
				Filename:  header.Filename,
				Mimetype:  mimetype,
				Location:  location,
				AssetType: constant.GetAssetType(mimetype),
			}

			res, err := uc.CreateMedia(r.Context(), req)
			if err != nil {
				response.Error(w, err)
				return
			}

			response.OK(w, res)
		})
	})

	return r
}
