package dataaccess

import (
	"context"
	"database/sql"
	"student_classes_management_service/pkg/application/interfaces"
	"student_classes_management_service/pkg/data-access/entity"

	"github.com/uptrace/bun"
)

type usersDA struct {
	dbc *bun.DB
}

func (s *usersDA) CreateUser(ctx context.Context, user *entity.User) (sql.Result, error) {
	return s.dbc.NewInsert().Model(user).Exec(ctx)
}

func (s *usersDA) GetUsers(c context.Context) ([]entity.User, error) {
	var list []entity.User
	err := s.dbc.NewSelect().Model(&list).Scan(c)

	return list, err
}

func NewUserDA(dbc *bun.DB) interfaces.UsersDA {
	return &usersDA{dbc}
}
