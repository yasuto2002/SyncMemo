package entity

import "encoding/json"

type Casual struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}

func (c *Casual) MarshalBinary() (data []byte, err error) {
	return json.Marshal(c)
}

func (c *Casual) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, c)
}
