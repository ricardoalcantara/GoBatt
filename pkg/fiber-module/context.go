package fibermodule

import "github.com/gofiber/fiber/v2"

type Context struct {
	ctx  *fiber.Ctx
	code int
}

func (c *Context) String(message string) {
	c.ctx.Status(c.code).SendString(message)
}

func (c *Context) Json(message any) {
	c.ctx.Status(c.code).JSON(message)
}

func (c *Context) Status(code int) {
	c.code = code
}
