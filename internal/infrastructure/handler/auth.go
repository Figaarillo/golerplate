package handler

import (
	"fmt"
	"net/http"

	"github.com/Figaarillo/golerplate/internal/application/usecase"
	"github.com/Figaarillo/golerplate/internal/domain/entity"
	"github.com/Figaarillo/golerplate/internal/domain/repository"
	"github.com/Figaarillo/golerplate/internal/shared/config"
	"github.com/Figaarillo/golerplate/internal/shared/constants"
	"github.com/Figaarillo/golerplate/internal/shared/utils"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	authUC      *usecase.AuthUseCase
	userUC      *usecase.UserUseCase
	respository repository.UserRepository
	validator   *validator.Validate
}

func NewAuthHandler(env *config.EnvVars, userRepo repository.UserRepository) *AuthHandler {
	return &AuthHandler{
		authUC:    usecase.NewAuthUseCase(env),
		userUC:    usecase.NewUserUseCase(userRepo),
		validator: validator.New(),
	}
}

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	var err error

	if err := utils.DecodeReqBody(r, &user); err != nil {
		utils.NewHTTPResponse(w).InternalServerError(err.Error(), nil)
	}

	if err := h.validator.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		validationErrors := make(map[string]string)
		for _, err := range errors {
			validationErrors[err.Field()] = err.Tag()
		}
		utils.NewHTTPResponse(w).InternalServerError(err.Error(), nil)
		return
	}

	if user, err = h.userUC.Create(user); err != nil {
		utils.NewHTTPResponse(w).Conflict(err.Error(), nil)
		return
	}

	token, err := h.authUC.GenerateTokens(user.ID.String())
	if err != nil {
		utils.NewHTTPResponse(w).InternalServerError(err.Error(), nil)
		return
	}

	utils.NewHTTPResponse(w).OK("User registered successfully", token)
}

func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	refreshToken, err := utils.GetHeader(r, constants.RefreshToken)
	if err != nil {
		utils.NewHTTPResponse(w).Unauthorized(err.Error(), nil)
		return
	}

	newAccessToken, err := h.authUC.RefreshAndStoreAccessToken(refreshToken)
	if err != nil {
		utils.NewHTTPResponse(w).Unauthorized(err.Error(), nil)
		return
	}

	utils.NewHTTPResponse(w).OK("New access token generated successfully", map[string]string{"access_token": newAccessToken})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var authUser entity.AuthUser
	if err := utils.DecodeReqBody(r, &authUser); err != nil {
		utils.NewHTTPResponse(w).InternalServerError(err.Error(), nil)
		return
	}

	if err := h.validator.Struct(authUser); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.NewHTTPResponse(w).InternalServerError(fmt.Sprintf("validation error: %s", errors), nil)
		return
	}

	user, err := h.userUC.GetByProp("email", authUser.Email)
	if err != nil {
		utils.NewHTTPResponse(w).NotFound(err.Error(), nil)
		return
	}

	if err := user.ComparePassword(authUser.Password); err != nil {
		utils.NewHTTPResponse(w).Unauthorized(err.Error(), nil)
		return
	}

	token, err := h.authUC.GenerateTokens(user.ID.String())
	if err != nil {
		utils.NewHTTPResponse(w).InternalServerError(err.Error(), nil)
		return
	}

	utils.NewHTTPResponse(w).OK("User logged in successfully", token)
}
