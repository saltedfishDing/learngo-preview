package repositories

import (
	"gorm.io/gorm"
	"preview/models"
	"preview/repositories/contract"
)

type RecordRepositorie struct {
	db *gorm.DB
}

func NewRecordRepositorie(db *gorm.DB) contract.RecordContract {
	return &RecordRepositorie{db: db}
}

func (r *RecordRepositorie) Find(record models.Record) (records []models.Record, total int64, err error) {
	db := r.db

	if record.OrderBy != "" {
		desc := " asc"
		if record.Desc != "" {
			desc = " " + record.Desc
		}
		db = db.Order(record.OrderBy + desc)
	}

	page := record.Page
	if page <= 0 {
		page = 1
	}

	limit := record.Limit
	if limit <= 0 {
		limit = 10
	}
	err = db.Debug().Offset(page - 1).Limit(limit).Find(&records).Error
	db.Model(record).Count(&total)
	return

}
