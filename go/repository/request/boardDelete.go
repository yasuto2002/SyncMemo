package request

type BoardDelete struct {
	BoardId string `json:"boardId" validate:"required"`
}
