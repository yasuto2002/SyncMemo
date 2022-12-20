package response

import "go.mongodb.org/mongo-driver/bson"

type BoardList struct {
	Boards []bson.M `json:"boards"`
}
