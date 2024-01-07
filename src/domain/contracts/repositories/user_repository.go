package repositories

import "goHexBoilerplate/src/domain/contracts/entities"

type GetByPropertiesParams struct {
	id    int
	name  string
	email string
}

type UserRepository interface {
	GetById(id string) (entities.User, error)
	GetByProperties(params GetByPropertiesParams) ([]entities.User, error)
}
