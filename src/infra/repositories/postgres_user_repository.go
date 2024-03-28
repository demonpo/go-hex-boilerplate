package repositories

import (
	"goHexBoilerplate/src/db"
	"goHexBoilerplate/src/domain/contracts/entities"
	"goHexBoilerplate/src/domain/contracts/repositories"
	entitiesInfra "goHexBoilerplate/src/infra/entities"
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
	newUser := entitiesInfra.User{
		Name: "Daniel", Email: "soul.danielssss@hotmail.com",
	}
	userRepository.db.DB.Model(&entitiesInfra.User{}).Create(&newUser)
	return entities.User{
		Id:    int(newUser.ID),
		Name:  newUser.Name,
		Email: newUser.Email,
	}, nil
}

func (userRepository *PostgresUserRepository) GetByProperties(params repositories.GetByPropertiesParams) ([]entities.User, error) {
	return []entities.User{entities.User{Id: 1, Name: "Daniel", Email: "soul.daniel@hotmail.com"}, entities.User{Id: 2, Name: "Xavier", Email: "xavicoGarcia12@gmail.com"}}, nil
}
