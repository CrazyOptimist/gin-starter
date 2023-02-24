package user

import (
	"gin-starter/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	UserService UserService
}

func NewUserController(db *gorm.DB) UserController {
	userRepository := NewUserRepository(db)
	userService := NewUserService(userRepository)
	return UserController{UserService: userService}
}

// FindAll godoc
// @Summary Retrieves users
// @Tags users
// @Success 200	{array} User
// @Failure 500 {object} utils.HTTPError
// @Router /admin/users [get]
func (u *UserController) FindAll(c *gin.Context) {
	users := u.UserService.FindAll()
	c.JSON(http.StatusOK, users)
}

// FindById godoc
// @Summary Retrieves a user by ID
// @Tags users
// @Param id path integer true "User ID"
// @Success 200	{object} User
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /admin/users/{id} [get]
func (u *UserController) FindById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.NewError(c, http.StatusBadRequest, err)
	}

	user, err := u.UserService.FindById(uint(id))
	if err != nil {
		utils.NewError(c, http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

// Create godoc
// @Summary Create a new user
// @Tags users
// @Param request body CreateUserDto true "CreateUserDto"
// @Success 201	{array} User
// @Failure 400 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /admin/users [post]
func (u *UserController) Create(c *gin.Context) {
	var dto CreateUserDto
	if err := c.BindJSON(&dto); err != nil {
		utils.NewError(c, http.StatusBadRequest, err)
	}

	user, err := u.UserService.Create(dto)
	if err != nil {
		utils.NewError(c, http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusCreated, user)
}

// Update godoc
// @Summary Update user
// @Tags users
// @Param id path integer true "User ID"
// @Param request body CreateUserDto true "UpdateUserDto"
// @Success 200	{array} User
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /admin/users/{id} [patch]
func (u *UserController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.NewError(c, http.StatusBadRequest, err)
	}

	if _, err := u.UserService.FindById(uint(id)); err != nil {
		utils.NewError(c, http.StatusNotFound, err)
	}

	var dto UpdateUserDto
	if err := c.BindJSON(&dto); err != nil {
		utils.NewError(c, http.StatusBadRequest, err)
	}

	user, err := u.UserService.Update(dto, uint(id))
	if err != nil {
		utils.NewError(c, http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, user)
}

// Delete godoc
// @Summary Delete user
// @Tags users
// @Param id path integer true "User ID"
// @Success 200
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Router /admin/users/{id} [delete]
func (u *UserController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.NewError(c, http.StatusBadRequest, err)
	}

	if err := u.UserService.Delete(uint(id)); err != nil {
		utils.NewError(c, http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, nil)
}
