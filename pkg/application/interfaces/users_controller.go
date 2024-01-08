package interfaces

import (
	"context"

	"github.com/labstack/echo/v4"
)

type UsersController interface {
	GetUsers(context.Context) error
	CreateUser(ctx echo.Context) error
}
