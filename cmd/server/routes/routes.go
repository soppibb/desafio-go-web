package routes

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/bootcamp-go/desafio-go-web/cmd/server/handler"
	"github.com/bootcamp-go/desafio-go-web/internal/domain"
	"github.com/bootcamp-go/desafio-go-web/internal/tickets"
	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
}

func (r *Router) Setup() error {
	err := r.SetTicketsRoutes()

	if err != nil {
		return err
	}

	return nil
}

func (r *Router) SetTicketsRoutes() error {

	allTickets, err := LoadTicketsFromFile("/Users/totorres/Desktop/desafio-go-web/cmd/server/tickets.csv")

	if err != nil {
		return err
	}

	repository := tickets.NewRepository(allTickets)
	service := tickets.NewService(repository)
	ticketHandler := handler.NewTicketHandler(service)

	ticket := r.router.Group("/tickets")
	{
		ticket.GET("/getByCountry/:name", ticketHandler.GetTotal())
		ticket.GET("/getAverage/:name", ticketHandler.GetAverageDestination())
	}

	return nil
}

func NewRouter(router *gin.Engine) *Router {
	return &Router{router: router}
}

func LoadTicketsFromFile(path string) ([]domain.Ticket, error) {

	var ticketList []domain.Ticket

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	csvR := csv.NewReader(file)
	data, err := csvR.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	for _, row := range data {
		price, err := strconv.ParseFloat(row[5], 64)
		if err != nil {
			return []domain.Ticket{}, err
		}
		ticketList = append(ticketList, domain.Ticket{
			Id:      row[0],
			Name:    row[1],
			Email:   row[2],
			Country: row[3],
			Time:    row[4],
			Price:   price,
		})
	}

	return ticketList, nil
}
