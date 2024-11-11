package repositories

import (
	"github.com/google/uuid"
	"goHexBoilerplate/src/modules/user/domain/contracts/entities"
)

type GetByPropertiesParams struct {
	id    uuid.UUID
	name  string
	email string
}

type Create struct {
	Name  string
	Email string
}

type UserRepository interface {
	GetById(id uuid.UUID) (*entities.User, error)
	GetByProperties(params GetByPropertiesParams) ([]entities.User, error)
	Create(params Create) (*entities.User, error)
}
