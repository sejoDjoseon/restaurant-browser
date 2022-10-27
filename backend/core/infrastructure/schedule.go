package restaurant

var mockRestaurantsOpened = map[string]bool{
	"111111": true,
	"222222": false,
}

type ScheduleService interface {
	isOpen(rstID string) bool
}

type service struct{}

func NewScheduleService() ScheduleService {
	return &service{}
}

// Mock function
func (s *service) isOpen(rstID string) bool {
	open, ok := mockRestaurantsOpened[rstID]
	return open && ok
}
