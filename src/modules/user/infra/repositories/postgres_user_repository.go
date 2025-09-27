package repositories

import (
	"fmt"
	"goHexBoilerplate/src/db"
	domainEntities "goHexBoilerplate/src/modules/user/domain/contracts/entities"
	domainRepositories "goHexBoilerplate/src/modules/user/domain/contracts/repositories"
	infraEntities "goHexBoilerplate/src/modules/user/infra/entities"

	"github.com/google/uuid"
)

type PostgresUserRepository struct {
	db *db.DB
}

func NewPostgresUserRepository(db *db.DB) *PostgresUserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}

func (userRepository *PostgresUserRepository) Create(
	params domainRepositories.Create,
) (*domainEntities.User, error) {
	gormUser := &infraEntities.User{
		Name:  params.Name,
		Email: params.Email,
	}
	if err := userRepository.db.DB.Create(gormUser).Error; err != nil {
		return nil, err
	}
	return &domainEntities.User{
		Id:        gormUser.ID,
		Name:      gormUser.Name,
		Email:     gormUser.Email,
		UpdatedAt: gormUser.UpdatedAt,
		CreatedAt: gormUser.CreatedAt,
	}, nil
}

func (userRepository *PostgresUserRepository) GetById(id uuid.UUID) (*domainEntities.User, error) {
	var gormUser infraEntities.User
	if err := userRepository.db.DB.First(&gormUser, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("failed to find user by ID %s: %w", id, err)
	}
	return &domainEntities.User{
		Id:        gormUser.ID,
		Name:      gormUser.Name,
		Email:     gormUser.Email,
		UpdatedAt: gormUser.UpdatedAt,
		CreatedAt: gormUser.CreatedAt,
	}, nil
}

func (userRepository *PostgresUserRepository) GetByProperties(
	params domainRepositories.GetByPropertiesParams,
) ([]domainEntities.User, error) {
	// Not implemented yet. Return demo data as before.
	return []domainEntities.User{
		{Id: uuid.UUID{}, Name: "Daniel", Email: "soul.daniel@hotmail.com"},
		{Id: uuid.UUID{}, Name: "Xavier", Email: "xavicoGarcia12@gmail.com"},
	}, nil
}
