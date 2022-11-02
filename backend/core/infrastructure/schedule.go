package restaurant

var mockRestaurantsOpened = map[string]bool{
	"111111": true,
	"222222": false,
	"333333": true,
	"444444": true,
	"555555": true,
	"666666": false,
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
