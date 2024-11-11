package repositories

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"goHexBoilerplate/ent/user"
	"goHexBoilerplate/src/db"
	"goHexBoilerplate/src/modules/user/domain/contracts/entities"
	"goHexBoilerplate/src/modules/user/domain/contracts/repositories"
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
	params repositories.Create,
) (*entities.User, error) {
	newUser, err := userRepository.db.DB.User.Create().
		SetName(params.Name).
		SetEmail(params.Email).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return &entities.User{
		Id:        newUser.ID,
		Name:      newUser.Name,
		Email:     newUser.Email,
		UpdatedAt: newUser.UpdatedAt,
		CreatedAt: newUser.CreatedAt,
	}, nil
}

func (userRepository *PostgresUserRepository) GetById(id uuid.UUID) (*entities.User, error) {
	foundUser, err := userRepository.db.DB.User.
		Query().
		Where(user.IDEQ(id)).
		Only(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to find user by ID %d: %w", id, err)
	}
	return &entities.User{
		Id:        foundUser.ID,
		Name:      foundUser.Name,
		Email:     foundUser.Email,
		UpdatedAt: foundUser.UpdatedAt,
		CreatedAt: foundUser.CreatedAt,
	}, nil
}

func (userRepository *PostgresUserRepository) GetByProperties(
	params repositories.GetByPropertiesParams,
) ([]entities.User, error) {
	return []entities.User{
		entities.User{Id: uuid.UUID{}, Name: "Daniel", Email: "soul.daniel@hotmail.com"},
		entities.User{Id: uuid.UUID{}, Name: "Xavier", Email: "xavicoGarcia12@gmail.com"},
	}, nil
}
