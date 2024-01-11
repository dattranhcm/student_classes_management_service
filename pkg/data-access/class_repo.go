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
	s.dbc.RegisterModel((*entity.StudentClass)(nil))
	return s.dbc.NewInsert().Model(class).Exec(ctx)
}

func (s *classRepo) GetClasses(c context.Context) ([]entity.Class, error) {
	var list []entity.Class
	err := s.dbc.NewSelect().Model(&list).Relation("Teacher").Scan(c)

	return list, err
}

func (c classRepo) GetById(ctx context.Context, id string, userId string) (*entity.Class, error) {
	class := new(entity.Class)
	count, err := c.dbc.NewSelect().Model(class).
		Where("class_id = ?", id).
		Where("teacher_id = ?", userId).
		Relation("Teacher").
		ScanAndCount(ctx)

	if count == 0 {
		return nil, err
	}

	return class, err
}

func (c *classRepo) AssignStudent(ctx context.Context, studentClass []entity.StudentClass) (sql.Result, error) {
	return c.dbc.NewInsert().Model(&studentClass).Exec(ctx)
}

func NewClassRepo(dbc *bun.DB) interfaces.ClassesRepository {
	dbc.RegisterModel((*entity.StudentClass)(nil))
	return &classRepo{dbc}
}
