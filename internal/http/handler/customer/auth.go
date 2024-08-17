package customer

import (
	"net/http"

	"github.com/ffajarpratama/gommerce-api/constant"
	"github.com/ffajarpratama/gommerce-api/internal/http/request"
	"github.com/ffajarpratama/gommerce-api/internal/http/response"
	"github.com/ffajarpratama/gommerce-api/lib/custom_validator"
	"github.com/ffajarpratama/gommerce-api/util"
	"github.com/google/uuid"
)

func (h *CustomerHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req request.Register
	err := custom_validator.ValidateStruct(r, &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	req.Role = constant.UserRoleCustomer

	res, err := h.uc.Register(r.Context(), &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, res)
}

func (h *CustomerHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req request.Login
	err := custom_validator.ValidateStruct(r, &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	req.Role = constant.UserRoleCustomer

	res, err := h.uc.Login(r.Context(), &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, res)
}

func (h *CustomerHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, _ := uuid.Parse(util.GetUserIDFromContext(ctx))

	res, err := h.uc.GetProfile(ctx, userID)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, res)
}
