package tickets

import (
	"context"

	"github.com/bootcamp-go/desafio-go-web/internal/domain"
)

type Service interface {
	GetTotalTickets(cxt context.Context, location string) (dt []domain.Ticket, err error)
	AverageDestination(cxt context.Context, location string) (avg float64, err error)
}
type service struct {
	repository Repository
}

func NewService(rep Repository) Service {
	return &service{
		repository: rep,
	}
}

func (s *service) GetTotalTickets(cxt context.Context, location string) (dt []domain.Ticket, err error) {
	dt, err = s.repository.GetTicketByDestination(cxt, location)
	return
}
func (s *service) AverageDestination(cxt context.Context, location string) (avg float64, err error) {
	tot, err3 := s.repository.GetAll(cxt)
	if err != nil {
		err = err3
		return
	}
	ttub, errx := s.repository.GetTicketByDestination(cxt, location)
	if errx != nil {
		err = err3
		return
	}
	avg = float64(len(ttub)) / float64(len(tot))

	return
}
