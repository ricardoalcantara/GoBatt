package home

import (
	"github.com/ricardoalcantara/GoBatt/internal/core"
)

type HomeModule struct {
}

func (home *HomeModule) Register(gobatt *core.GoBatt) {
	gobatt.Register(NewHomeController)
	gobatt.Register(NewHomeService)
}

func (home *HomeModule) Init(gobatt *core.GoBatt) {
	gobatt.Invoke(func(h *HomeController) {})
}
