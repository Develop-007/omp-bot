package seat

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Seat {
	return allEntities
}

func (s *Service) Get(idx int) (*Seat, error) {
	return &allEntities[idx], nil
}
