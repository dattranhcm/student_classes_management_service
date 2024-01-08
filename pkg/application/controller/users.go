package controller

import (
	"fmt"
	"net/http"
	"student_classes_management_service/pkg/application/interfaces"
	"student_classes_management_service/pkg/application/model"
	"github.com/labstack/echo/v4"
)

type userController struct {
	userService interfaces.UsersService
}

func (api *userController) CreateUser(e echo.Context) error {
	registerInfo := new(model.Users)
	student, err := api.userService.CreateUser(e.Request().Context(), *registerInfo)
	if err != nil {
		log.Print(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not register new student")
	}

	return e.JSON(http.StatusCreated, student)
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
