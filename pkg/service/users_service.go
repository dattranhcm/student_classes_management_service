package service

import (
	"context"
	"student_classes_management_service/pkg/application/interfaces"
	"student_classes_management_service/pkg/application/model"
	"student_classes_management_service/pkg/data-access/entity"
	"student_classes_management_service/pkg/utils"
)

type usersService struct {
	db interfaces.UsersRepository
}

func (s *usersService) CreateUser(ctx context.Context, info model.Users) (model.Users, error) {
	newUser := entity.User{
		Username: info.Username,
		UserType: info.UserType,
		FullName: info.FullName,
		Password: utils.HashPassword(info.Password),
	}

	_, err := s.db.CreateUser(ctx, &newUser)
	return mapToUserModel(newUser), err
}

func (s *usersService) GetUsers(c context.Context) ([]model.Users, error) {
	list, err := s.db.GetUsers(c)
	if err != nil {
		return nil, err
	}

	result := make([]model.Users, len(list))
	for i, v := range list {
		result[i] = mapToUserModel(v)
	}
	return result, nil
}

func (s *usersService) FindByUsername(ctx context.Context, username string) (model.Users, error) {
	student, err := s.db.FindByUsername(ctx, username)

	return mapToUserModel(student), err
}

func (s *usersService) FindByUsernameToLogin(ctx context.Context, username string) (model.Users, error) {
	student, err := s.db.FindByUsername(ctx, username)

	return getUserSensitiveInfo(student), err
}


func (s *usersService) GetClasses(ctx context.Context, username string) ([]model.Class, error) {
	classDtos, err := s.db.GetClasses(ctx, username)
	result := make([]model.Class, len(classDtos))

	for i, v := range classDtos {
		result[i] = mapToClassModel(v)
	}

	return result, err
}

func mapToUserModel(s entity.User) model.Users {
	classes := make([]model.Class, len(s.Classes))

	for i, v := range s.Classes {
		classes[i] = mapToClassModel(v)
	}

	return model.Users{
		UserId:   s.UserId,
		Username: s.Username,
		UserType: s.UserType,
		FullName: s.FullName,
		Classes: classes,
	}
	
}

func getUserSensitiveInfo(s entity.User) model.Users {
	return model.Users{
		UserId:   s.UserId,
		Username: s.Username,
		UserType: s.UserType,
		FullName: s.FullName,
		Password: s.Password,
	}
	
}


func NewUserService(db interfaces.UsersRepository) *usersService {
	return &usersService{db}
}
