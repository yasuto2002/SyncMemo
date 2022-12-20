package request

type Make struct {
	Name     string `validate:"required"`
	Password string `validate:"required"`
}
