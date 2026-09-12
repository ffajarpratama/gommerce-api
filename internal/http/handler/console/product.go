package console

import (
	"net/http"

	"github.com/ffajarpratama/gommerce-api/internal/http/request"
	"github.com/ffajarpratama/gommerce-api/internal/http/response"
	"github.com/ffajarpratama/gommerce-api/lib/custom_validator"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (h *ConsoleHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req request.CreateProduct
	err := custom_validator.ValidateStruct(r, &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	err = h.uc.CreateProduct(r.Context(), &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, nil)
}

func (h *ConsoleHandler) FindAndCountProduct(w http.ResponseWriter, r *http.Request) {
	var params request.ListProductQuery
	params.BaseQuery = request.NewBaseQuery(r)
	params.CategoryName = r.URL.Query().Get("category_name")

	res, cnt, err := h.uc.FindAndCountProduct(r.Context(), &params)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Paging(w, res, params.Page, params.Limit, cnt)
}

func (h *ConsoleHandler) FindOneProduct(w http.ResponseWriter, r *http.Request) {
	productID, _ := uuid.Parse(chi.URLParam(r, "product_id"))
	res, err := h.uc.FindOneProduct(r.Context(), productID)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, res)
}

func (h *ConsoleHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var req request.UpdateProduct
	err := custom_validator.ValidateStruct(r, &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	req.ProductID, _ = uuid.Parse(chi.URLParam(r, "product_id"))

	err = h.uc.UpdateProduct(r.Context(), &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, nil)
}

func (h *ConsoleHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID, _ := uuid.Parse(chi.URLParam(r, "product_id"))
	err := h.uc.DeleteProduct(r.Context(), productID)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, nil)
}
