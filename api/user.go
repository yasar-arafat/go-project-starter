package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	userController "github.com/yasar-arafat/go-project-starter/controller/user"
	"github.com/yasar-arafat/go-project-starter/model"
	"github.com/yasar-arafat/go-project-starter/utils"
)

// SignUp godoc
// @Summary Register a new user
// @Description Register a new user
// @ID sign-up
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body userRegisterRequest true "User info for registration"
// @Success 201 {object} userResponse
// @Failure 400 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Router /users [post]
func SignUp(c echo.Context) error { // TODO validation remaining
	var u model.User
	req := &userRegisterRequest{}
	if err := req.bind(c, &u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if err := userController.CreateUser(&u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusCreated, newUserResponse(&u))
}

// GetAllUser godoc
// @Summary Get all user
// @Description Gets all Registered users
// @Tags user
// @Produce  json
// @Success 200 {object} userResponse
// @Failure 400 {object} utils.Error
// @Failure 401 {object} utils.Error
// @Failure 422 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Router /users [get]
func GetAllUser(c echo.Context) error {
	u, err := userController.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	c.Logger().Info("Get All User =", c.Response().Header().Get(echo.HeaderXRequestID))
	return c.JSON(http.StatusOK, getAllUsersResponse(u))
}
