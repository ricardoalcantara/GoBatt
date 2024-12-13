package ginmodule

import "github.com/gin-gonic/gin"

func NewGinModule() *gin.Engine {
	r := gin.Default()
	return r
}
