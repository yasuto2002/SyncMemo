package request

type Memo struct {
	Id       string `json:"id" validate:"required"`
	Text     string `json:"text" validate:"required"`
	X        int    `json:"x" validate:"required"`
	Y        int    `json:"y" validate:"required"`
	ActionId int    `json:"actionId" validate:"required"`
	BoardId  string `json:"boardId" validate:"required"`
}
