package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"preview/services/contract"
)

type recordWebService struct {
}

func NewRecordWebService() contract.RecordWebContract {
	return &recordWebService{}
}

func (rw *recordWebService) List(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"ip":    c.ClientIP(),
		"title": "你的IP",
	})
}
