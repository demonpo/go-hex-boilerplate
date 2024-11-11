package services

import (
	"github.com/google/uuid"
	"goHexBoilerplate/src/modules/user/domain/contracts/entities"
	"goHexBoilerplate/src/modules/user/domain/contracts/repositories"
)

type UserService struct {
	userRepository repositories.UserRepository
}

type CreateInput struct {
	Name  string
	Email string
}

func NewUserService(userRepository repositories.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (userService *UserService) GetById(id uuid.UUID) (*entities.User, error) {
	userRepository := userService.userRepository
	return userRepository.GetById(id)
}

func (userService *UserService) Create(input CreateInput) (*entities.User, error) {
	userRepository := userService.userRepository
	return userRepository.Create(repositories.Create{Name: input.Name, Email: input.Email})
}
