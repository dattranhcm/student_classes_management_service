package service

import (
	"context"
	"student_classes_management_service/pkg/application/interfaces"
	"student_classes_management_service/pkg/application/model"
	"student_classes_management_service/pkg/data-access/entity"
)

type classesService struct {
	db interfaces.ClassesRepository
}

func (s *classesService) CreateClass(ctx context.Context, info model.Class) (model.Class, error) {
	newClass := entity.Class{
		ClassName: info.ClassName,
		TeacherId: info.TeacherId,
		DayOfWeek: info.DayOfWeek,
		StartTime: info.StartTime,
		EndTime:   info.EndTime,
	}

	_, err := s.db.CreateClass(ctx, &newClass)
	return mapToClassModel(newClass), err
}

func (s *classesService) GetClasses(c context.Context) ([]model.Class, error) {
	list, err := s.db.GetClasses(c)
	if err != nil {
		return nil, err
	}

	result := make([]model.Class, len(list))
	for i, v := range list {
		result[i] = mapToClassModel(v)
	}
	return result, nil
}

func mapToClassModel(s entity.Class) model.Class {

	return model.Class{
		ClassId:     s.ClassId,
		ClassName:   s.ClassName,
		TeacherId:   s.TeacherId,
		TeacherName: s.Teacher.Username,
		DayOfWeek:   s.DayOfWeek,
		StartTime:   s.StartTime,
		EndTime:     s.EndTime,
		CreatedAt:   s.CreatedAt,
		UpdatedAt:   s.UpdatedAt,
	}
}

func NewClassService(db interfaces.ClassesRepository) *classesService {
	return &classesService{db}
}
