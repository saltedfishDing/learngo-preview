package contract

import "github.com/gin-gonic/gin"

type RecordWebContract interface {
	List(c *gin.Context)
}
