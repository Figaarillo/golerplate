package handler

import (
	"net/http"

	"github.com/Figaarillo/golerplate/internal/application/usecase"
	"github.com/Figaarillo/golerplate/internal/domain/entity"
	"github.com/Figaarillo/golerplate/internal/domain/repository"
	"github.com/Figaarillo/golerplate/internal/shared/config"
	"github.com/Figaarillo/golerplate/internal/shared/utils"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	authUC      *usecase.AuthUseCase
	userUC      *usecase.UserUseCase
	respository repository.UserRepository
	validator   *validator.Validate
}

func NewAuthHandler(env *config.EnvVars, authRepo repository.AuthRepository, userRepo repository.UserRepository) *AuthHandler {
	return &AuthHandler{
		authUC:    usecase.NewAuthUseCase(env, authRepo),
		userUC:    usecase.NewUserUseCase(userRepo),
		validator: validator.New(),
	}
}

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	var err error

	if err := utils.DecodeReqBody(r, &user); err != nil {
		utils.HandleHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	if err := h.validator.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		validationErrors := make(map[string]string)
		for _, err := range errors {
			validationErrors[err.Field()] = err.Tag()
		}
		utils.HandleHTTPError(w, errors, http.StatusUnprocessableEntity)
		return
	}

	if user, err = h.userUC.Create(user); err != nil {
		utils.HandleHTTPError(w, err, http.StatusConflict)
		return
	}

	token, err := h.authUC.GenerateAndStoreTokens(user.ID)
	if err != nil {
		utils.HandleHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	utils.HandleHTTPResponse(w, "User registered successfully", http.StatusCreated, token)
}

func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	refreshToken := utils.GetHeader(r, "refresh-token")

	newAccessToken, err := h.authUC.RefreshAndStoreAccessToken(refreshToken)
	if err != nil {
		utils.HandleHTTPError(w, err, http.StatusUnauthorized)
		return
	}
	utils.HandleHTTPResponse(w, "New access token generated successfully", http.StatusOK, map[string]string{
		"access_token": newAccessToken,
	})
}
