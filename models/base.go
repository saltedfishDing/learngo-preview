package models

type Base struct {
	ID      uint   `gorm:"primarykey" json:"id"`
	Page    int    `gorm:"-" json:"-" form:"page"`
	Limit   int    `gorm:"-" json:"-" form:"limit"`
	OrderBy string `gorm:"-" json:"-" form:"order_by"`
	Desc    string `gorm:"-" json:"-" form:"desc"`
}

type StringArrayJson []string
