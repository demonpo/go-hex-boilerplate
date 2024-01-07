package services

import (
	"goHexBoilerplate/src/domain/contracts/entities"
	"goHexBoilerplate/src/domain/contracts/repositories"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (userService *UserService) GetById(id string) (entities.User, error) {
	userRepository := userService.userRepository
	return userRepository.GetById(id)
}
