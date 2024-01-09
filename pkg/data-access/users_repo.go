package dataaccess

import (
	"context"
	"database/sql"
	"student_classes_management_service/pkg/application/interfaces"
	"student_classes_management_service/pkg/data-access/entity"

	"github.com/uptrace/bun"
)

type userRepo struct {
	dbc *bun.DB
}

func (s *userRepo) CreateUser(ctx context.Context, user *entity.User) (sql.Result, error) {
	return s.dbc.NewInsert().Model(user).Exec(ctx)
}

func (s *userRepo) GetUsers(c context.Context) ([]entity.User, error) {
	var list []entity.User
	err := s.dbc.NewSelect().Model(&list).Scan(c)

	return list, err
}

func NewUserRepo(dbc *bun.DB) interfaces.UsersRepository {
	return &userRepo{dbc}
}
