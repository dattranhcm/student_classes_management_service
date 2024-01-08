package interfaces

import (
	"context"
	"database/sql"
	"student_classes_management_service/pkg/data-access/entity"
)

type UsersDA interface {
	CreateUser(ctx context.Context, user *entity.User) (sql.Result, error)
	GetUsers(ctx context.Context) ([]entity.User, error)
}
