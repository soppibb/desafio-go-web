package tickets

import "context"

type Service interface {
	GetTotalTickets(ctx context.Context, s string) (int, error)
	AverageDestination(ctx context.Context, s string) (float64, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return service{repo}
}

func (sv service) GetTotalTickets(ctx context.Context, s string) (int, error) {
	tickets, err := sv.repo.GetTicketByDestination(ctx, s)
	if err != nil {
		return 0, err
	}
	return len(tickets), nil
}

func (sv service) AverageDestination(ctx context.Context, s string) (float64, error) {
	tickets, err := sv.repo.GetTicketByDestination(ctx, s)
	if err != nil {
		return 0, err
	}
	ticketsall, err := sv.repo.GetAll(ctx)
	if err != nil {
		return 0, err
	}
	return float64(len(tickets)) / float64(len(ticketsall)), nil
}
