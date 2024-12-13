package core

import (
	"github.com/ricardoalcantara/GoBatt/internal/logger"
	"github.com/rs/zerolog/log"
	"go.uber.org/dig"
)

func NewGoBatt() GoBatt {
	container := dig.New()
	container.Provide(logger.NewLogger)
	container.Provide(func() *GoBatt {
		return &GoBatt{}
	})

	goBatt := GoBatt{
		container: container,
	}

	return goBatt
}

type GoBatt struct {
	container *dig.Container
	logger    logger.ILogger
	engine    IEngine
}

func (g *GoBatt) Engine() IEngine {
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
	var engine IEngine
	var _logger logger.ILogger
	err := g.container.Invoke(func(e IEngine, l logger.ILogger) {
		engine = e
		_logger = l
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}

	g.engine = engine
	g.logger = _logger

	g.engine.Run()
}
