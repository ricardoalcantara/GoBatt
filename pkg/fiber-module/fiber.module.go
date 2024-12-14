package fibermodule

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ricardoalcantara/GoBatt/pkg/core"
)

func NewFiberModule() core.IEngine {
	app := fiber.New()
	return &Engine{engine: app}
}
