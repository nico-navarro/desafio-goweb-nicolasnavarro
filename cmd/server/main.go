package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nico-navarro/desafio-goweb-nicolasnavarro/cmd/server/handler"
	"github.com/nico-navarro/desafio-goweb-nicolasnavarro/internal/domain"
	"github.com/nico-navarro/desafio-goweb-nicolasnavarro/internal/tickets"
)

func main() {

	// Cargo csv.
	dbTickets, err := LoadTicketsFromFile("tickets.csv")
	if err != nil {
		panic("Couldn't load tickets")
	}

	ticketRepository := tickets.NewRepository(dbTickets)
	ticketService := tickets.NewService(ticketRepository)
	ticketController := handler.NewTicketController(ticketService)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

	ticketRouter := r.Group("/ticket")
	ticketRouter.GET("/getByCountry/:dest", ticketController.GetTicketsByCountry())
	ticketRouter.GET("/getAverage/:dest", ticketController.AverageDestination())

	if err := r.Run(); err != nil {
		panic(err)
	}

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
