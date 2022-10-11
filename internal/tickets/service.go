package tickets

type Service interface {
	GetTotalTickets(destination string) (int, error)
	AverageDestination(destination string) (int, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetTotalTickets(destination string) (int, error) {
	tickets, err := s.repository.GetTicketByDestination(destination)
	if err != nil {
		return 0, err
	}
	totalTickets := len(tickets)

	return totalTickets, nil
}

func (s *service) AverageDestination(destination string) (int, error) {
	return 0, nil
}
