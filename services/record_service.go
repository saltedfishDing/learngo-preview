package services

import (
	"preview/models"
	"preview/repositories"
	repositories_contract "preview/repositories/contract"
	"preview/services/contract"
)

type recordService struct {
	repositorie repositories_contract.RecordContract
}

func NewRecordService() contract.RecordContract {
	return &recordService{
		repositorie: repositories.NewRecordRepositorie(models.DB),
	}
}

func (r *recordService) Get(record models.Record) (records []models.Record, total int64, err error) {
	records, total, err = r.repositorie.Find(record)
	return
}
