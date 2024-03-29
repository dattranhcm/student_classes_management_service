package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"student_classes_management_service/pkg/application/interfaces"
	"student_classes_management_service/pkg/application/model"
	"student_classes_management_service/pkg/utils"

	"github.com/labstack/echo/v4"
)

type classController struct {
	classesService interfaces.ClassesService
}

func (api *classController) CreateClass(e echo.Context) error {
	claims, _ := utils.GetTokenClaims(e)
	newClass := new(model.Class)
	if err := utils.BindAndValidate(e, newClass); err != nil {
		return err
	}
	teacherId, err := strconv.Atoi(claims.ID)
	if err != nil {
		return err
	}
	user, err := api.classesService.CreateClass(e.Request().Context(), *newClass, teacherId)
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

func (c *classController) GetById(e echo.Context) error {
	classId := e.Param("class_id")
	claims, err := utils.GetTokenClaims(e)

	if err != nil {
		return err
	}

	result, err := c.classesService.GetById(e.Request().Context(), classId, claims.ID)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, result)
}

func (c *classController) AssignStudent(e echo.Context) error {
	classId := e.Param("class_id")

	claims, err := utils.GetTokenClaims(e)
	if err != nil {
		return err
	}

	assignStudent := new(model.AssignStudent)
	if err := utils.BindAndValidate(e, assignStudent); err != nil {
		return echo.ErrBadRequest
	}

	if err := c.classesService.AssignStudent(e.Request().Context(), classId, assignStudent.StudentIds, claims.ID); err != nil {
		return err
	}

	class, _ := c.classesService.GetById(e.Request().Context(), classId, claims.ID)
	return e.JSON(http.StatusOK, class)
}

func NewClassController(u interfaces.ClassesService) interfaces.ClassesController {
	return &classController{u}
}
