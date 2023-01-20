package entity

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Memo struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	BoardId string             `json:"boardId"`
	Text    string             `json:"text"`
	X       int                `json:"x"`
	Y       int                `json:"y"`
}

func (m *Memo) MarshalBinary() (data []byte, err error) {
	return json.Marshal(m)
}

func (m *Memo) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}
