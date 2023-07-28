package contract

import "preview/models"

type RecordContract interface {
	Get(record models.Record) ([]models.Record, int64, error)
}
