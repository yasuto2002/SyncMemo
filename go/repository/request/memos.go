package request

type Memos struct {
	Id string `json:"id" validate:"required"`
}
