package interfaces

import (
	"context"
	"database/sql"
	"student_classes_management_service/pkg/application/model"
	"student_classes_management_service/pkg/data-access/entity"

	"github.com/labstack/echo/v4"
)

type (
	UsersRepository interface {
		CreateUser(ctx context.Context, user *entity.User) (sql.Result, error)
		GetUsers(ctx context.Context) ([]entity.User, error)
	}

	UsersService interface {
		CreateUser(ctx context.Context, info model.Users) (model.Users, error)
		GetUsers(context.Context) ([]model.Users, error)
	}

	UsersController interface {
		GetUsers(ctx echo.Context) error
		CreateUser(ctx echo.Context) error
	}
)