package fibermodule

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ricardoalcantara/GoBatt/pkg/core"
)

type Engine struct {
	engine *fiber.App
}

func (g Engine) Start() {
	g.engine.Listen(":8080")
}

func (g Engine) Get(path string, handler func(core.IContext)) {
	g.engine.Get(path, func(ctx *fiber.Ctx) error {
		handler(&Context{ctx: ctx, code: http.StatusOK})
		return nil
	})
}
func (g Engine) Post(path string, handler func(core.IContext)) {
	g.engine.Post(path, func(ctx *fiber.Ctx) error {
		handler(&Context{ctx: ctx, code: http.StatusCreated})
		return nil
	})
}
func (g Engine) Put(path string, handler func(core.IContext)) {
	g.engine.Put(path, func(ctx *fiber.Ctx) error {
		handler(&Context{ctx: ctx, code: http.StatusAccepted})
		return nil
	})
}
func (g Engine) Delete(path string, handler func(core.IContext)) {
	g.engine.Delete(path, func(ctx *fiber.Ctx) error {
		handler(&Context{ctx: ctx, code: http.StatusNoContent})
		return nil
	})
}
func (g Engine) Patch(path string, handler func(core.IContext)) {
	g.engine.Patch(path, func(ctx *fiber.Ctx) error {
		handler(&Context{ctx: ctx, code: http.StatusAccepted})
		return nil
	})
}
