package repositories

import (
	"goHexBoilerplate/src/domain/contracts/entities"
	"goHexBoilerplate/src/domain/contracts/repositories"
	"goHexBoilerplate/src/infra/db"
)

type PostgresUserRepository struct {
	db *db.DB
}

func NewPostgresUserRepository(db *db.DB) *PostgresUserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}

func (userRepository *PostgresUserRepository) GetById(id string) (entities.User, error) {
	userRepository.db.DB.Create(&entities.User{Name: "Daniel", Email: ""})
	return entities.User{Id: 1, Name: "Daniel", Email: "soul.daniel@hotmail.com"}, nil
}

func (userRepository *PostgresUserRepository) GetByProperties(params repositories.GetByPropertiesParams) ([]entities.User, error) {
	return []entities.User{entities.User{Id: 1, Name: "Daniel", Email: "soul.daniel@hotmail.com"}, entities.User{Id: 2, Name: "Xavier", Email: "xavicoGarcia12@gmail.com"}}, nil
}
