package routers

import (
	v1 "preview/api/v1"

	"github.com/gin-gonic/gin"
)

func initRecordWeb(r *gin.Engine) {
	v := r.Group("/recordWeb")
	recordWeb := v1.NewRecordWeb()
	v.GET("", recordWeb.List)
}
