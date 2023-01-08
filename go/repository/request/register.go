package request

type Register struct {
	Token string `json:"token" validate:"required"`
	Mail  string `json:"mail" validate:"required"`
}
