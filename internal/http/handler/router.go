package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ffajarpratama/gommerce-api/config"
	"github.com/ffajarpratama/gommerce-api/constant"
	console_handler "github.com/ffajarpratama/gommerce-api/internal/http/handler/console"
	customer_handler "github.com/ffajarpratama/gommerce-api/internal/http/handler/customer"
	"github.com/ffajarpratama/gommerce-api/internal/http/middleware"
	"github.com/ffajarpratama/gommerce-api/internal/http/response"
	"github.com/ffajarpratama/gommerce-api/internal/usecase"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

func NewHTTPRouter(cnf *config.Config, uc usecase.IFaceUsecase) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recover)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	}))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response.JsonResponse{
			Error: &response.ErrorResponse{
				Code:    constant.DefaultNotFoundError,
				Status:  http.StatusNotFound,
				Message: "please check url",
			},
		})
	})

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response.JsonResponse{
			Error: &response.ErrorResponse{
				Code:    constant.DefaultMethodNotAllowedError,
				Status:  http.StatusMethodNotAllowed,
				Message: "method not allowed",
			},
		})
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response.JsonResponse{
			Success: true,
			Data: map[string]interface{}{
				"app-name": "ffajarpratama/gommerce-api",
			},
		})
	})

	r.Route("/api", func(api chi.Router) {
		api.Route("/v1", func(v1 chi.Router) {
			v1.Mount("/console", console_handler.NewHandler(cnf, uc))
			v1.Mount("/customer", customer_handler.NewHandler(cnf, uc))
		})
	})

	return r
}
