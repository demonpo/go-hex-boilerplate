package repositories

import (
	"goHexBoilerplate/src/domain/contracts/entities"
	"goHexBoilerplate/src/domain/contracts/repositories"
)

type PostgresUserRepository struct {
}

func NewPostgresUserRepository() *PostgresUserRepository {
	return &PostgresUserRepository{}
}

func (userRepository *PostgresUserRepository) getById(id int) (entities.User, error) {
	return entities.User{}, nil
}

func (userRepository *PostgresUserRepository) getByProperties(params repositories.GetByPropertiesParams) ([]entities.User, error) {
	return []entities.User{entities.User{Id: 1, Name: "Daniel", Email: "soul.daniel@hotmail.com"}, entities.User{Id: 2, Name: "Xavier", Email: "xavicoGarcia12@gmail.com"}}, nil
}
