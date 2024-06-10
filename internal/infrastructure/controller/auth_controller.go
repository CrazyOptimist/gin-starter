package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gin-starter/internal/domain/auth"
	"gin-starter/internal/infrastructure/helper"
	"gin-starter/internal/infrastructure/repository"
	"gin-starter/pkg/utils"
)

type authController struct {
	AuthService auth.AuthService
}

func NewAuthController(db *gorm.DB) *authController {
	userRepository := repository.NewUserRepository(db)
	authHelper := helper.NewAuthHelper()
	authService := auth.NewAuthService(userRepository, authHelper)
	return &authController{AuthService: *authService}
}

// Register godoc
// @Summary Register new user
// @Tags auth
// @Param request body auth.RegisterDto true "RegisterDto"
// @Success 201	{object} auth.LoginResponse
// @Failure 400 {object} utils.HttpError
// @Failure 500 {object} utils.HttpError
// @Router /auth/register [post]
func (a *authController) Register(c *gin.Context) {
	var registerDto auth.RegisterDto
	if err := c.BindJSON(&registerDto); err != nil {
		utils.RaiseHttpError(c, http.StatusBadRequest, err)
		return
	}

	loginResponse, err := a.AuthService.Register(&registerDto)
	if err != nil {
		utils.RaiseHttpError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, loginResponse)
}

// Login godoc
// @Summary Login user
// @Tags auth
// @Param request body auth.LoginDto true "LoginDto"
// @Success 201	{object} auth.LoginResponse
// @Failure 400 {object} utils.HttpError
// @Failure 401 {object} utils.HttpError
// @Failure 404 {object} utils.HttpError
// @Router /auth/login [post]
func (a *authController) Login(c *gin.Context) {
	var loginDto auth.LoginDto
	if err := c.BindJSON(&loginDto); err != nil {
		utils.RaiseHttpError(c, http.StatusBadRequest, err)
		return
	}

	loginResponse, err := a.AuthService.Login(&loginDto)
	if err != nil {
		utils.RaiseHttpError(c, http.StatusUnauthorized, err)
		return
	}

	c.JSON(http.StatusCreated, loginResponse)
}

// Logout godoc
// @Summary Logout user (Invalidates refresh token)
// @Tags auth
// @Success 200
// @Failure 401 {object} utils.HttpError
// @Failure 500 {object} utils.HttpError
// @Router /auth/logout [post]
// @Security JWT
func (a *authController) Logout(c *gin.Context) {
	var logoutDto auth.LogoutDto
	if err := c.BindJSON(&logoutDto); err != nil {
		utils.RaiseHttpError(c, http.StatusBadRequest, err)
		return
	}

	if err := a.AuthService.Logout(&logoutDto); err != nil {
		utils.RaiseHttpError(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

// RefreshToken godoc
// @Summary Refresh tokens and invalidate the previous refresh token
// @Tags auth
// @Param request body auth.LogoutDto true "TokenRefresh DTO"
// @Success 201	{object} auth.LoginResponse
// @Failure 400 {object} utils.HttpError
// @Failure 500 {object} utils.HttpError
// @Router /auth/refresh [post]
func (a *authController) RefreshToken(c *gin.Context) {
	var refreshDto auth.LogoutDto
	if err := c.BindJSON(&refreshDto); err != nil {
		utils.RaiseHttpError(c, http.StatusBadRequest, err)
		return
	}

	refreshResponse, err := a.AuthService.RefreshToken(&refreshDto)
	if err != nil {
		if errors.Is(err, auth.ErrTokenBlacklisted) {
			utils.RaiseHttpError(c, http.StatusUnauthorized, err)
			return
		}
		utils.RaiseHttpError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, refreshResponse)
}