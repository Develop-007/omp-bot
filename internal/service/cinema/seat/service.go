package seat

type SeatService interface {
	Describe(seat_id uint64) (*Seat, error)
	List(cursor uint64, limit uint64) []Seat
	Create(Seat) (uint64, error)
	Update(seat_id uint64, seat Seat) error
	Remove(seat_id uint64) (bool, error)
}

type DummySeatService struct{}

func NewService() *DummySeatService {
	return &DummySeatService{}
}

func (s *DummySeatService) Describe(idx uint64) (*Seat, error) {
	return &allEntities[idx], nil
}

func (s *DummySeatService) List(cursor uint64, limit uint64) []Seat {
	return allEntities
}

func (s *DummySeatService) Create(seat Seat) (uint64, error) {
	allEntities = append(allEntities, seat)
	return uint64(len(allEntities) - 1), nil
}

func (s *DummySeatService) Update(seat Seat) error {
	return nil
}

func (s *DummySeatService) Remove(idx int) (bool, error) {
	allEntities = append(allEntities[:idx], allEntities[idx+1:]...)
	return true, nil
}
