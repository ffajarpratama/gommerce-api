package console

import (
	"net/http"

	"github.com/ffajarpratama/gommerce-api/config"
	"github.com/ffajarpratama/gommerce-api/internal/http/middleware"
	"github.com/ffajarpratama/gommerce-api/internal/usecase"
	"github.com/go-chi/chi/v5"
)

type ConsoleHandler struct {
	cnf *config.Config
	uc  usecase.IFaceUsecase
}

func NewHTTPHandler(cnf *config.Config, uc usecase.IFaceUsecase) http.Handler {
	r := chi.NewRouter()
	h := ConsoleHandler{
		cnf: cnf,
		uc:  uc,
	}

	r.Route("/auth", func(auth chi.Router) {
		auth.Post("/login", h.Login)

		auth.Route("/profile", func(profile chi.Router) {
			profile.Use(middleware.Authorize(cnf.JWT.Secret))
			profile.Get("/", h.GetProfile)
		})
	})

	r.Group(func(private chi.Router) {
		private.Use(middleware.Authorize(cnf.JWT.Secret))

		private.Route("/product", func(product chi.Router) {
			product.Post("/", h.CreateProduct)
			product.Get("/", h.FindAndCountProduct)
			product.Get("/{product_id}", h.FindOneProduct)
			product.Put("/{product_id}", h.UpdateProduct)
			product.Delete("/{product_id}", h.DeleteProduct)
		})
	})

	return r
}
