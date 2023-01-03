package request

type Make struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"passwowd"`
}
