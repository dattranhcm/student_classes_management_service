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

func (api *userController) Login(e echo.Context) error {
	credential := new(model.LoginCredential)
	if err := utils.BindAndValidate(e, credential); err != nil {
		return err
	}
	student, err := api.userService.FindByUsername(e.Request().Context(), credential.Username)
	fmt.Println("AAA")
	fmt.Println(student)
	if err != nil {
		fmt.Println(err)
	//	return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credential")
	}

	if err := utils.ComparePassword(student.Password, credential.Password); err == nil {
		token := utils.GenerateToken(student)
		return e.JSON(http.StatusOK, map[string]string{"token": token})
	}

	return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credential")
}

func (api *userController) GetClasses(e echo.Context) error {
	claims, _ := utils.GetTokenClaims(e)
	classes, err := api.userService.GetClasses(e.Request().Context(), claims.Username)

	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, classes)
}


func (api *userController) Profile(e echo.Context) error {
	claims, err := utils.GetTokenClaims(e)

	if err != nil {
		return err
	}

	student, err := api.userService.FindByUsername(e.Request().Context(), claims.Username)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Fail to parse into from token")
	}

	return e.JSON(http.StatusOK, student)
}

func NewUserController(u interfaces.UsersService) interfaces.UsersController {
	return &userController{u}
}
