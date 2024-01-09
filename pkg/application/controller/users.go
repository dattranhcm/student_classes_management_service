package controller

import (
	"fmt"
	"log"
	"net/http"
	"student_classes_management_service/pkg/application/interfaces"
	"student_classes_management_service/pkg/application/model"
	"student_classes_management_service/pkg/utils"

	"github.com/labstack/echo/v4"
)

type userController struct {
	userService interfaces.UsersService
}

func (api *userController) CreateUser(e echo.Context) error {
	registerInfo := new(model.Users)
	if err := utils.BindAndValidate(e, registerInfo); err != nil {
		return err
	}

	user, err := api.userService.CreateUser(e.Request().Context(), *registerInfo)
	if err != nil {
		log.Print(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not register new User")
	}

	return e.JSON(http.StatusCreated, user)
}

func (api *userController) GetUsers(c echo.Context) error {
	users, err := api.userService.GetUsers(c.Request().Context())
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong")
	}

	return c.JSON(http.StatusOK, users)
}

func NewUserController(u interfaces.UsersService) interfaces.UsersController {
	return &userController{u}
}
