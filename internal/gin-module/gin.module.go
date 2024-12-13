package ginmodule

import (
	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/GoBatt/internal/core"
)

func NewGinModule() core.IEngine {
	r := gin.Default()
	return &Engine{engine: r}
}

type Engine struct {
	engine *gin.Engine
}

func (g Engine) Run() {
	g.engine.Run(":8080")
}

func (g Engine) GET(path string, handler func(core.IContext)) {
	g.engine.GET(path, func(ctx *gin.Context) {
		handler(&Context{ctx: ctx})
	})
}

type Context struct {
	ctx *gin.Context
}

func (c *Context) String(code int, message string) {
	c.ctx.String(code, message)
}
