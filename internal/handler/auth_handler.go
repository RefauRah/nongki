package handler

import (
	"encoding/json"
	"net/http"
	"nongki/internal/request"
	"nongki/internal/response"
	"nongki/internal/usecase"
	"nongki/pkg/constant"
	"nongki/pkg/helpers"

	validator "github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
	validate    *validator.Validate
}

func NewAuthHandler(authUsecase usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		authUsecase: authUsecase,
		validate:    validator.New()}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req request.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.ErrorResponse(constant.INVALID_REQUEST))
		return
	}

	if err := h.validate.Struct(req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.ErrorResponse(helpers.FormatValidationError(err)))
		return
	}

	data, token, err := h.authUsecase.Register(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.ErrorResponse(err.Error()))
		return
	}

	var registerResponse response.RegisterResponse
	registerResponse.UserDomainToRegisterResponse(data, token)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.SuccessResponse(constant.REGISTER_SUCCESS, registerResponse))
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req request.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.ErrorResponse(constant.INVALID_REQUEST))
		return
	}

	if err := h.validate.Struct(req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.ErrorResponse(helpers.FormatValidationError(err)))
		return
	}

	data, token, err := h.authUsecase.Login(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.ErrorResponse(err.Error()))
		return
	}

	var loginResponse response.LoginResponse
	loginResponse.UserDomainToLoginResponse(data, token)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.SuccessResponse(constant.LOGIN_SUCCESS, loginResponse))
}
