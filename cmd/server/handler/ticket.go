package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nico-navarro/desafio-goweb-nicolasnavarro/internal/tickets"
)

type TicketController struct {
	service tickets.Service
}

func NewTicketController(s tickets.Service) *TicketController {
	return &TicketController{
		service: s,
	}
}

func (c *TicketController) GetTicketsByCountry() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		destination := ctx.Param("dest")

		tickets, err := c.service.GetTotalTickets(destination)
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		ctx.JSON(200, tickets)
	}
}

func (c *TicketController) AverageDestination() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		destination := ctx.Param("dest")

		avg, err := c.service.AverageDestination(destination)
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		ctx.JSON(200, avg)
	}
}
