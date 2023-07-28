package v1

import (
	"github.com/gin-gonic/gin"
	"preview/services"
	"preview/services/contract"
)

type recordWeb struct {
	service contract.RecordWebContract
}

func NewRecordWeb() *recordWeb {
	return &recordWeb{
		service: services.NewRecordWebService(),
	}
}

func (rw *recordWeb) List(c *gin.Context) {
	rw.service.List(c)
}
