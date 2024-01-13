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
		FindByUsername(ctx context.Context, username string) (entity.User, error)
		GetClasses(ctx context.Context, username string) ([]entity.Class, error)
	}

	UsersService interface {
		CreateUser(ctx context.Context, info model.Users) (model.Users, error)
		GetUsers(context.Context) ([]model.Users, error)
		FindByUsername(ctx context.Context, username string) (model.Users, error)
		GetClasses(context.Context, string) ([]model.Class, error)
	}

	UsersController interface {
		Login(ctx echo.Context) error
		GetUsers(ctx echo.Context) error
		CreateUser(ctx echo.Context) error
		Profile(ctx echo.Context) error
		GetClasses(ctx echo.Context) error
	}
)
