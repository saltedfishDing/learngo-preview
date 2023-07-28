package routers

import (
	v1 "preview/api/v1"

	"github.com/gin-gonic/gin"
)

func initRecord(r *gin.Engine) {
	v := r.Group("/record")
	record := v1.NewRecord()
	v.GET("", record.Get)
}
