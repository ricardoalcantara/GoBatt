package ginmodule

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/GoBatt/pkg/core"
)

type Engine struct {
	engine *gin.Engine
}

func (g Engine) Start() {
	g.engine.Run(":8080")
}

func (g Engine) Get(path string, handler func(core.IContext)) {
	g.engine.GET(path, func(ctx *gin.Context) {
		handler(&Context{ctx: ctx, code: http.StatusOK})
	})
}
func (g Engine) Post(path string, handler func(core.IContext)) {
	g.engine.POST(path, func(ctx *gin.Context) {
		handler(&Context{ctx: ctx, code: http.StatusCreated})
	})
}
func (g Engine) Put(path string, handler func(core.IContext)) {
	g.engine.PUT(path, func(ctx *gin.Context) {
		handler(&Context{ctx: ctx, code: http.StatusAccepted})
	})
}
func (g Engine) Delete(path string, handler func(core.IContext)) {
	g.engine.DELETE(path, func(ctx *gin.Context) {
		handler(&Context{ctx: ctx, code: http.StatusNoContent})
	})
}
func (g Engine) Patch(path string, handler func(core.IContext)) {
	g.engine.PATCH(path, func(ctx *gin.Context) {
		handler(&Context{ctx: ctx, code: http.StatusAccepted})
	})
}
