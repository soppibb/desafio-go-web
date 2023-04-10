package handler

import (
	"net/http"

	"github.com/bootcamp-go/desafio-go-web/internal/tickets"
	"github.com/gin-gonic/gin"
)

type TicketHandler struct {
	service tickets.Service
}

func NewTicketHandler(service tickets.Service) *TicketHandler {
	return &TicketHandler{service: service}
}

func (t *TicketHandler) GetTotal() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		name := ctx.Param("name")

		total, err := t.service.GetTotalTickets(ctx, name)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, nil)
		}

		ctx.JSON(http.StatusAccepted, gin.H{"data": total})

	}
}

func (t *TicketHandler) GetAverageDestination() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		name := ctx.Param("name")

		avg, err := t.service.AverageDestination(ctx, name)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, nil)
		}

		ctx.JSON(http.StatusAccepted, gin.H{"data": avg})
	}
}

// type Service struct {
// 	service tickets.Service
// }

// func NewService(s tickets.Service) *Service {
// 	return &Service{
// 		service: s,
// 	}
// }

// func (s *Service) GetTicketsByCountry() gin.HandlerFunc {
// 	return func(c *gin.Context) {

// 		destination := c.Param("dest")

// 		tickets, err := s.service.GetTotalTickets(c, destination)
// 		if err != nil {
// 			c.String(http.StatusInternalServerError, err.Error(), nil)
// 			return
// 		}

// 		c.JSON(200, tickets)
// 	}
// }

// func (s *Service) AverageDestination() gin.HandlerFunc {
// 	return func(c *gin.Context) {

// 		destination := c.Param("dest")

// 		avg, err := s.service.AverageDestination(c, destination)
// 		if err != nil {
// 			c.String(http.StatusInternalServerError, err.Error(), nil)
// 			return
// 		}

// 		c.JSON(200, avg)
// 	}
// }
