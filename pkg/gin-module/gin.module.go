package ginmodule

import (
	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/GoBatt/pkg/core"
)

func NewGinModule() core.IEngine {
	r := gin.Default()
	return &Engine{engine: r}
}
