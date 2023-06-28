package repositories

import "goHexBoilerplate/src/domain/contracts/entities"

type GetByPropertiesParams struct {
	id    int
	name  string
	email string
}

type UserRepository interface {
	getById(id int) entities.User
	getByProperties(params GetByPropertiesParams) []entities.User
}
