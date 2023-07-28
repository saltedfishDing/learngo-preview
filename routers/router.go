package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 设置web的静态目录和html目录
	r.Static("/assets", "./web/static")
	r.LoadHTMLGlob("./web/view/*")
	// 上传下载目录
	//r.StaticFS("/upload", http.Dir("./web/upload"))
	initRecord(r)
	initRecordWeb(r)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"ip":    c.ClientIP(),
			"title": "你的IP",
		})
	})
	return r
}
