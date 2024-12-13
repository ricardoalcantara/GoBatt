package main

import (
	"github.com/ricardoalcantara/GoBatt/internal/core"
	fibermodule "github.com/ricardoalcantara/GoBatt/internal/fiber-module"
	"github.com/ricardoalcantara/GoBatt/internal/home"
)

func main() {
	goBatt := core.NewGoBatt()
	// goBatt.Register(ginmodule.NewGinModule)
	goBatt.Register(fibermodule.NewFiberModule)
	goBatt.AddModule(&home.HomeModule{})

	goBatt.Start()
}
