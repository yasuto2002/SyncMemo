package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Memo struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	BoardId string             `json:"boardId"`
	Text    string             `json:"text"`
	X       int                `json:"X"`
	Y       int                `json:"Y"`
}
