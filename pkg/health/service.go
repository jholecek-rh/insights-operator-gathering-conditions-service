package health

type ServiceInterface interface{}

type Service struct{}

func NewService() *Service {
	return &Service{}
}
