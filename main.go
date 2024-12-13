package main

import (
	"github.com/ricardoalcantara/GoBatt/internal/core"
	"github.com/ricardoalcantara/GoBatt/internal/home"
)

func main() {
	goBatt := core.NewGoBatt()
	goBatt.AddModule(&home.HomeModule{})

	goBatt.Start()
}
