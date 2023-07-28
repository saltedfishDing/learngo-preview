package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"preview/models"
	"preview/pkg/response"
	"preview/services"
	"preview/services/contract"
)

type record struct {
	service contract.RecordContract
}

func NewRecord() *record {
	return &record{
		service: services.NewRecordService(),
	}
}

func (r *record) Get(c *gin.Context) {
	var (
		params models.Record
	)

	if err := c.BindQuery(&params); err != nil {
		fmt.Println(err)
		response.FailBindingError(err, c)
		return
	}

	result, total, err := r.service.Get(params)

	if err != nil {
		response.Fail(c)
		return
	}

	resultData := response.ResultData{
		"list":  result,
		"total": total,
	}

	response.OkWithData(resultData, c)

}
