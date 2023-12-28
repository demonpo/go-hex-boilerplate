package services

import "goHexBoilerplate/src/domain/contracts/entities"

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (userService *UserService) GetById(id string) (entities.User, error) {
	return entities.User{
		Id:    0,
		Name:  "",
		Email: "",
	}, nil
}
