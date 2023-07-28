package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Record struct {
	Base
	UserId    string          `gorm:"column:user_id" json:"user_id"`
	LoginTime int64           `gorm:"column:login_time" json:"login_time"`
	LoginDate string          `gorm:"column:login_date" json:"login_date"`
	Images    StringArrayJson `gorm:"column:images" json:"images"`
}

func (*Record) TableName() string {
	return "record"
}

func (sa StringArrayJson) Value() (driver.Value, error) {
	return json.Marshal(sa)
}

func (sa *StringArrayJson) Scan(value interface{}) error {
	fmt.Println(value)
	if len(value.([]byte)) > 0 {
		return json.Unmarshal(value.([]byte), &sa)
	} else {
		return nil
	}

}
