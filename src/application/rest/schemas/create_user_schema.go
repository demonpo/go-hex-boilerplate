package schemas

type CreateUserSchema struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}
