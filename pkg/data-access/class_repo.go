package dataaccess

import (
	"context"
	"database/sql"
	"student_classes_management_service/pkg/application/interfaces"
	"student_classes_management_service/pkg/data-access/entity"

	"github.com/uptrace/bun"
)

type classRepo struct {
	dbc *bun.DB
}

func (s *classRepo) CreateClass(ctx context.Context, class *entity.Class) (sql.Result, error) {
	return s.dbc.NewInsert().Model(class).Exec(ctx)
}

func (s *classRepo) GetClasses(c context.Context) ([]entity.Class, error) {
	var list []entity.Class
	err := s.dbc.NewSelect().Model(&list).Relation("Teacher").Scan(c)

	return list, err
}

func NewClassRepo(dbc *bun.DB) interfaces.ClassesRepository {
	return &classRepo{dbc}
}
