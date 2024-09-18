package handler

import (
	"encoding/json"
	"net/http"
	"nongki/internal/request"
	"nongki/internal/response"
	"nongki/internal/usecase"
	"nongki/pkg/constant"
	"nongki/pkg/helpers"
	middleware "nongki/pkg/midleware"

	validator "github.com/go-playground/validator/v10"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
	validate    *validator.Validate
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: userUsecase,
		validate: validator.New()}
}

func (h *UserHandler) GetMe(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserContextKey("userID")).(string)

	data, err := h.userUsecase.GetMe(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.ErrorResponse(err.Error()))
		return
	}

	var userResponse response.UserResponse
	userResponse.UserDomainToUserResponse(data)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.SuccessResponse(constant.GET_ME_SUCCESS, userResponse))
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserContextKey("userID")).(string)

	var userUpdateRequest request.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&userUpdateRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.ErrorResponse(constant.INVALID_REQUEST))
		return
	}

	if err := h.validate.Struct(userUpdateRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.ErrorResponse(helpers.FormatValidationError(err)))
		return
	}

	data, err := h.userUsecase.UpdateUser(userID, userUpdateRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response.InternalServerErrorResponse(constant.FAILED_UPDATE_PROFILE))
		return
	}

	var userResponse response.UserResponse
	userResponse.UserDomainToUserResponse(data)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.SuccessResponse(constant.SUCCESS_UPDATE_PROFILE, userResponse))
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserContextKey("userID")).(string)
	deletedBy := userID

	err := h.userUsecase.DeleteUser(userID, deletedBy)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response.InternalServerErrorResponse(constant.DELETED_FAILED))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.SuccessResponse(constant.DELETED_SUCCESS, nil))
}
