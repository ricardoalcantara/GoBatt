package home

type HomeService struct {
}

func NewHomeService() *HomeService {
	return &HomeService{}
}

func (s *HomeService) Index() string {
	return "hello world"
}
