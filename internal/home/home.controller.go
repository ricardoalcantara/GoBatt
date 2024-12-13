package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/GoBatt/internal/logger"
)

type HomeController struct {
	logger      logger.ILogger
	homeService *HomeService
}

func NewHomeController(r *gin.Engine, logger logger.ILogger, homeService *HomeService) *HomeController {
	controller := HomeController{
		logger:      logger,
		homeService: homeService,
	}
	r.GET("/", controller.Index)

	return &controller
}

func (c *HomeController) Index(ctx *gin.Context) {
	c.logger.Info("Index")
	index := c.homeService.Index()
	ctx.String(http.StatusOK, index)
}
