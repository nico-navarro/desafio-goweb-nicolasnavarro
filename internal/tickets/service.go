package tickets

import (
	"github.com/nico-navarro/desafio-goweb-nicolasnavarro/internal/domain"
)

type Service interface {
	GetTotalTickets(destination string) ([]domain.Ticket, error)
	AverageDestination(destination string) ([]domain.Ticket, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetTotalTickets(destination string) ([]domain.Ticket, error) {
	return []domain.Ticket{}, nil
}

func (s *service) AverageDestination(destination string) ([]domain.Ticket, error) {
	return []domain.Ticket{}, nil
}
