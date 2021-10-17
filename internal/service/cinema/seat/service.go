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

func (s *Service) Add(seat Seat) {
	allEntities = append(allEntities, seat)
}

func (s *Service) Delete(idx int) {
	allEntities = append(allEntities[:idx], allEntities[idx+1:]...)
}
