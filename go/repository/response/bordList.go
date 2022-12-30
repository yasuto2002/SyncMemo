package response

type BoardList struct {
	Id      string `bson:"_id" json:"id"`
	Name    string `bson:"name" json:"name"`
	Mail    string `bson:"mail" json:"Mail"`
	Paaword string `bson:"paaword" json:"paaword"`
}
