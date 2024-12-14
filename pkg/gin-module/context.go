package ginmodule

import "github.com/gin-gonic/gin"

type Context struct {
	ctx  *gin.Context
	code int
}

func (c *Context) String(message string) {
	c.ctx.String(c.code, message)
}

func (c *Context) Json(message any) {
	c.ctx.JSON(c.code, message)
}

func (c *Context) Status(code int) {
	c.code = code
}
