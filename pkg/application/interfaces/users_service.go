package interfaces

import (
	"context"
	"student_classes_management_service/pkg/application/model"
)

type UsersService interface {
	CreateUser(context.Context) (model.Users, error)
	GetUsers(context.Context) ([]model.Users, error)
}
