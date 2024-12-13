package core

import (
	"github.com/gin-gonic/gin"
	ginmodule "github.com/ricardoalcantara/GoBatt/internal/gin-module"
	"github.com/ricardoalcantara/GoBatt/internal/logger"
	"github.com/rs/zerolog/log"
	"go.uber.org/dig"
)

func NewGoBatt() GoBatt {
	container := dig.New()
	container.Provide(logger.NewLogger)
	container.Provide(ginmodule.NewGinModule)
	container.Provide(func() *GoBatt {
		return &GoBatt{}
	})

	var engine *gin.Engine
	var _logger logger.ILogger
	err := container.Invoke(func(r *gin.Engine, l logger.ILogger) {
		engine = r
		_logger = l
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}

	goBatt := GoBatt{
		container: container,
		logger:    _logger,
		engine:    engine,
	}

	return goBatt
}

type GoBatt struct {
	container *dig.Container
	logger    logger.ILogger
	engine    *gin.Engine
}

func (g *GoBatt) Engine() *gin.Engine {
	return g.engine
}

func (g *GoBatt) Container() *dig.Container {
	return g.container
}

func (g *GoBatt) Invoke(function interface{}) {
	g.container.Invoke(function)
}

func (g *GoBatt) Register(module interface{}) {
	g.container.Provide(module)
}

func (g *GoBatt) AddModule(module IModule) {
	module.Register(g)
	module.Init(g)
}

func (g *GoBatt) Start() {
	g.engine.Run()
}

type IModule interface {
	Register(*GoBatt)
	Init(*GoBatt)
}
