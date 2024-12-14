package home

import (
	"github.com/ricardoalcantara/GoBatt/pkg/core"
	"github.com/ricardoalcantara/GoBatt/pkg/logger"
)

type HomeController struct {
	logger      logger.ILogger
	homeService *HomeService
}

func NewHomeController(engine core.IEngine, logger logger.ILogger, homeService *HomeService) *HomeController {
	controller := HomeController{
		logger:      logger,
		homeService: homeService,
	}
	engine.Get("/", controller.Index)
	engine.Get("/json", controller.Json)

	return &controller
}

func (c *HomeController) Index(ctx core.IContext) {
	c.logger.Info("Index")
	index := c.homeService.Index()
	ctx.String(index)
}

func (c *HomeController) Json(ctx core.IContext) {
	c.logger.Info("Json")
	json := c.homeService.Json()
	ctx.Json(json)
}
