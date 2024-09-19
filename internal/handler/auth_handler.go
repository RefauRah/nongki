package handler

import (
	"encoding/json"
	"net/http"
	"nongki/internal/request"
	"nongki/internal/response"
	"nongki/internal/usecase"
	"nongki/pkg/constant"
	"nongki/pkg/helpers"
	"nongki/pkg/jwt"
	middleware "nongki/pkg/midleware"

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

func (h *AuthHandler) RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	var req request.RefreshTokenRequest
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

	userID := r.Context().Value(middleware.UserContextKey("userID")).(string)
	valid, err := jwt.ValidateRefreshToken(userID, req.RefreshToken)
	if !valid || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.ErrorResponse(constant.INVALID_TOKEN))
		return
	}

	newAccessToken, err := jwt.GenerateAccessToken(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.ErrorResponse(constant.GENERATE_TOKEN_FAILED))
		return
	}

	var refreshTokenResponse response.RefreshTokenResponse
	refreshTokenResponse.UserDomainToRefreshTokenResponse(newAccessToken)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.SuccessResponse(constant.GENERATE_TOKEN_SUCCESS, refreshTokenResponse))
}
