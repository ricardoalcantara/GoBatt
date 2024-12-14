package home

import "github.com/ricardoalcantara/GoBatt/internal/home/dto"

type HomeService struct {
}

func NewHomeService() *HomeService {
	return &HomeService{}
}

func (s *HomeService) Index() string {
	return "hello world"
}

func (s *HomeService) Json() dto.DefaultDTO {
	return dto.DefaultDTO{
		Message: "hello world",
	}
}
