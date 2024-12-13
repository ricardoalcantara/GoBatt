package fibermodule

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ricardoalcantara/GoBatt/internal/core"
)

func NewFiberModule() core.IEngine {
	app := fiber.New()
	return &Engine{engine: app}
}

type Engine struct {
	engine *fiber.App
}

func (g Engine) Run() {
	g.engine.Listen(":8080")
}

func (g Engine) GET(path string, handler func(core.IContext)) {
	g.engine.Get(path, func(ctx *fiber.Ctx) error {
		handler(&Context{ctx: ctx})
		return nil
	})
}

type Context struct {
	ctx *fiber.Ctx
}

func (c *Context) String(code int, message string) {
	c.ctx.Status(code).SendString(message)
}
