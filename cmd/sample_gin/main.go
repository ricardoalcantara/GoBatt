package main

import (
	"github.com/ricardoalcantara/GoBatt/internal/home"
	"github.com/ricardoalcantara/GoBatt/pkg/core"
	ginmodule "github.com/ricardoalcantara/GoBatt/pkg/gin-module"
	"github.com/ricardoalcantara/GoBatt/pkg/logger"
)

func main() {
	goBatt := core.NewGoBatt()
	goBatt.Register(logger.NewLogger)
	goBatt.Register(ginmodule.NewGinModule)
	goBatt.AddModule(&home.HomeModule{})

	goBatt.Start()
}
