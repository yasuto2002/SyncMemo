package response

type Make struct {
	ID string `validate:"required" json:"id"`
}
