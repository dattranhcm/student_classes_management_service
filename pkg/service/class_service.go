package service

import (
	"context"
	"log"
	"net/http"
	"student_classes_management_service/pkg/application/interfaces"
	"student_classes_management_service/pkg/application/model"
	"student_classes_management_service/pkg/data-access/entity"

	"github.com/labstack/echo/v4"
)

type classesService struct {
	db interfaces.ClassesRepository
}

func (s *classesService) CreateClass(ctx context.Context, info model.Class) (model.Class, error) {
	newClass := entity.Class{
		ClassName: info.ClassName,
		TeacherId: 1,
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

func (s *classesService) GetById(e context.Context, id string, userId string) (*model.Class, error) {
	classEntity, err := s.db.GetById(e, id, userId)

	if classEntity == nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, "Class not found")
	}

	result := mapToClassModel(*classEntity)
	return &result, err
}

func (s *classesService) AssignStudent(e context.Context, classId string, studentIds []string, userId string) error {
	_, err := s.db.GetById(e, classId, userId)
	if err != nil {
		return err
	}

	studentClass := make([]entity.StudentClass, len(studentIds))
	for i, v := range studentIds {
		studentClass[i] = entity.StudentClass{
			StudentID: v,
			ClassID:   classId,
		}
	}

	_, err = s.db.AssignStudent(e, studentClass)
	if err != nil {
		log.Print(err)
		return echo.ErrInternalServerError
	}

	return nil
}

func mapToClassModel(s entity.Class) model.Class {
	var students []model.Users
	for _, item := range s.Students {
		students = append(students, mapToUserModel(item))
	}
	var teacher *model.Users

	if s.Teacher != nil {
		value := mapToUserModel(*s.Teacher)
		teacher = &value
	} else {
		teacher = nil
	}

	return model.Class{
		ClassId:   s.ClassId,
		ClassName: s.ClassName,
		/*TeacherId: s.TeacherId,
			TeacherName: s.Teacher.Username,
		TeacherName: "Dat",  */
		DayOfWeek:   s.DayOfWeek,
		StartTime:   s.StartTime,
		EndTime:     s.EndTime,
		CreatedAt:   s.CreatedAt,
		UpdatedAt:   s.UpdatedAt,
		Teacher: teacher,
		Students: students,
	}
}

func NewClassService(db interfaces.ClassesRepository) *classesService {
	return &classesService{db}
}
