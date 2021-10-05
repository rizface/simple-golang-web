package web

type UserRequest struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}
