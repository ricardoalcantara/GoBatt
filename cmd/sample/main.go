package main

import (
	"github.com/ricardoalcantara/GoBatt/internal/home"
	"github.com/ricardoalcantara/GoBatt/pkg/core"
	fibermodule "github.com/ricardoalcantara/GoBatt/pkg/fiber-module"
	"github.com/ricardoalcantara/GoBatt/pkg/logger"
)

func main() {
	goBatt := core.NewGoBatt()
	goBatt.Register(logger.NewLogger)
	goBatt.Register(fibermodule.NewFiberModule)
	goBatt.AddModule(&home.HomeModule{})

	goBatt.Start()
}
