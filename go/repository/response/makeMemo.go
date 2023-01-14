package response

type MakeMemo struct {
	MemoID string `validate:"required" json:"memoId"`
}
