package contract

import "preview/models"

type RecordContract interface {
	Find(record models.Record) ([]models.Record, int64, error)
}
