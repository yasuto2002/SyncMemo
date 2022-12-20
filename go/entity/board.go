package entity

type Board struct {
	NAME     string
	MAIL     string
	PASSWORD string
}
type GetBoard struct {
	ID       map[string]string
	NAME     map[string]string
	MAIL     map[string]string
	PASSWORD map[string]string
}
