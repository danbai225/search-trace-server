package modle

import (
	"database/sql/driver"
	"encoding/json"
)

type Test struct {
	ID   int64
	Name string
	Age  int64
	Data Data `gorm:"type:json"`
}
type Data struct {
	UserInfo UserInfo `json:"user_info"`
	DataSize int64    `json:"data_size"`
	Comment  string   `json:"comment"`
}

func (c Data) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *Data) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

type UserInfo struct {
	Site        string `json:"site"`
	PhoneNumber string `json:"phone_number"`
	HeadAddress string `json:"head_address"`
	Distance    int64  `json:"distance"`
}
