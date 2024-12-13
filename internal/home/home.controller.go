package home

import (
	"net/http"

	"github.com/ricardoalcantara/GoBatt/internal/core"
	"github.com/ricardoalcantara/GoBatt/internal/logger"
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
	engine.GET("/", controller.Index)

	return &controller
}

func (c *HomeController) Index(ctx core.IContext) {
	c.logger.Info("Index")
	index := c.homeService.Index()
	ctx.String(http.StatusOK, index)
}
