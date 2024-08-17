package usecase

import (
	"context"
	"net/http"

	"github.com/ffajarpratama/gommerce-api/internal/http/request"
	"github.com/ffajarpratama/gommerce-api/internal/model"
	"github.com/ffajarpratama/gommerce-api/internal/repository"
	"github.com/ffajarpratama/gommerce-api/lib/custom_error"
	"github.com/ffajarpratama/gommerce-api/lib/hash"
	"github.com/ffajarpratama/gommerce-api/lib/jwt"
	"github.com/ffajarpratama/gommerce-api/util"
	"github.com/google/uuid"
)

// Register implements IFaceUsecase.
func (u *Usecase) Register(ctx context.Context, req *request.Register) (*model.User, error) {
	pwd, err := hash.HashAndSalt([]byte(req.Password))
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Name:        req.Name,
		Email:       req.Email,
		PhoneNumber: util.AddPhoneCode(req.PhoneNumber),
		Password:    pwd,
		Role:        req.Role,
	}

	err = u.repo.CreateUser(ctx, user, u.db)
	if err != nil {
		if repository.IsDuplicateErr(err) {
			err = custom_error.SetCustomError(&custom_error.ErrorContext{
				HTTPCode: http.StatusConflict,
				Message:  "email atau no. hp sudah digunakan",
			})

			return nil, err
		}

		return nil, err
	}

	claims := &jwt.CustomClaims{
		UserID: user.UserID.String(),
		Role:   string(user.Role),
	}

	user.AccessToken, err = jwt.GenerateToken(claims, u.cnf.JWT.Secret)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Login implements IFaceUsecase.
func (u *Usecase) Login(ctx context.Context, req *request.Login) (*model.User, error) {
	res, err := u.repo.FindOneUser(ctx, "email = ? AND role = ?", req.Email, req.Role)
	if err != nil {
		if repository.IsRecordNotfound(err) {
			err = custom_error.SetCustomError(&custom_error.ErrorContext{
				HTTPCode: http.StatusNotFound,
				Message:  "email atau password salah",
			})
		}

		return nil, err
	}

	err = hash.Compare([]byte(res.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	claims := &jwt.CustomClaims{
		UserID: res.UserID.String(),
		Role:   string(res.Role),
	}

	res.AccessToken, err = jwt.GenerateToken(claims, u.cnf.JWT.Secret)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetProfile implements IFaceUsecase.
func (u *Usecase) GetProfile(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	return u.repo.FindOneUser(ctx, "user_id = ?", userID)
}
