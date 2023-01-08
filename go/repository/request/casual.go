package request

type CasualUser struct {
	Name     string `json:"name" validate:"required"`
	Mail     string `json:"mail" validate:"required"`
	Password string `json:"password" validate:"required"`
}
