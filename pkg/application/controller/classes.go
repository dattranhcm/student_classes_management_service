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

type classController struct {
	classesService interfaces.ClassesService
}

func (api *classController) CreateClass(e echo.Context) error {
	newClass := new(model.Class)
	if err := utils.BindAndValidate(e, newClass); err != nil {
		return err
	}

	user, err := api.classesService.CreateClass(e.Request().Context(), *newClass)
	if err != nil {
		log.Print(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not create new Class")
	}

	return e.JSON(http.StatusCreated, user)
}

func (api *classController) GetClasses(c echo.Context) error {
	users, err := api.classesService.GetClasses(c.Request().Context())
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong")
	}

	return c.JSON(http.StatusOK, users)
}

func NewClassController(u interfaces.ClassesService) interfaces.ClassesController {
	return &classController{u}
}
