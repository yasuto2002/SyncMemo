package response

import "syncmemo/entity"

type ResMemos struct {
	Memos []entity.Memo `json:"memos" validate:"required"`
}
