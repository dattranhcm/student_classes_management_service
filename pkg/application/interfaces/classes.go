package interfaces

import (
	"context"
	"database/sql"
	"student_classes_management_service/pkg/application/model"
	"student_classes_management_service/pkg/data-access/entity"

	"github.com/labstack/echo/v4"
)

type (
	ClassesRepository interface {
		CreateClass(ctx context.Context, class *entity.Class) (sql.Result, error)
		GetClasses(ctx context.Context) ([]entity.Class, error)
		GetById(ctx context.Context, id string, userId string) (*entity.Class, error)
		AssignStudent(ctx context.Context, students []entity.StudentClass) (sql.Result, error)
	}

	ClassesService interface {
		CreateClass(ctx context.Context, info model.Class, teacherId int) (model.Class, error)
		GetClasses(context.Context) ([]model.Class, error)
		GetById(ctx context.Context, id string, userId string) (*model.Class, error)
		AssignStudent(ctx context.Context, classId string, studentIds []string, userId string) error
	}

	ClassesController interface {
		CreateClass(ctx echo.Context) error
		GetClasses(ctx echo.Context) error
		GetById(e echo.Context) error
		AssignStudent(e echo.Context) error
	}
)
