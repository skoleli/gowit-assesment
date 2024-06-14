package main

import (
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/go-playground/validator/v10"
)

const sortModeTime = "time"
const sortModeStatus = "status"

// Changed Ticket struct's field types and added a validation logic
type Ticket struct {
	ID          int
	Title       string `validate:"required,min=1"`
	Description string
	Status      string `validate:"required,min=1"`
	CreatedAt   time.Time
}

// implemented builder design pattern to avoid long ticket creation lines and simplify the logic
type TicketBuilder struct {
	nextId    int
	validator *validator.Validate
}

func NewTicketBuilder() *TicketBuilder {
	return &TicketBuilder{
		nextId:    1,
		validator: validator.New(),
	}
}

func (tb *TicketBuilder) NewTicket(title string, description string, status string) (Ticket, error) {
	ticket := Ticket{
		ID:          tb.nextId,
		Title:       title,
		Description: description,
		Status:      status,
		CreatedAt:   time.Now(),
	}
	tb.nextId++

	err := tb.validator.Struct(ticket)
	if err != nil {
		tb.nextId--
		return Ticket{}, fmt.Errorf("validation failure: %w", err)
	}
	return ticket, nil
}

/*
 * commented this function out since it's no longer in use
 * since CreatedAt field's type is time.Time,
 * but it definetly
 * needed an error handling logic. so i added it.
 */
// func parseTimeInUTC(timeStr, timezone string) (time.Time, error) {
// 	loc, err := time.LoadLocation(timezone)
// 	if err != nil {
// 		loc = time.UTC
// 	}
// 	localTime, err := time.ParseInLocation(time.RFC3339, timeStr, loc)
// 	return localTime.UTC(), err
// }

func BadSortTickets(tickets []Ticket, mode string) []Ticket {
	if mode == sortModeStatus {
		return sortTicketsByStatus(tickets)
	} else { // shortened this since default sort is also by time.
		return sortTicketsByTime(tickets)
	}
}

func sortTicketsByTime(tickets []Ticket) []Ticket {
	sort.SliceStable(tickets, func(i, j int) bool {
		return tickets[i].CreatedAt.After(tickets[j].CreatedAt)
	})
	return tickets
}

func sortTicketsByStatus(tickets []Ticket) []Ticket {
	sort.SliceStable(tickets, func(i, j int) bool {
		if tickets[i].Status == tickets[j].Status {
			return tickets[i].CreatedAt.After(tickets[j].CreatedAt)
		}
		return tickets[i].Status < tickets[j].Status
	})
	return tickets
}

func PrintTickets(tickets []Ticket) {
	for _, ticket := range tickets {
		fmt.Printf("%+v\n", ticket)
	}
}

func main() {
	ticketBuilder := NewTicketBuilder()
	tickets := []Ticket{}
	ticketValues := [][]string{
		{"Ticket 1", "First ticket", "closed"},
		{"Ticket 2", "Second ticket", "open"},
		{"Ticket 3", "Third ticket", "in-progress"},
	}

	for _, value := range ticketValues {
		newTicket, err := ticketBuilder.NewTicket(value[0], value[1], value[2])
		if err != nil {
			log.Printf("Failed to build ticket (%s, %s, %s) | %v\n", value[0], value[1], value[2], err)
		} else {
			tickets = append(tickets, newTicket)
		}
	}

	sortedByTime := BadSortTickets(tickets, sortModeTime)
	fmt.Println("Sorted by Time:")
	PrintTickets(sortedByTime)

	sortedByStatus := BadSortTickets(tickets, sortModeStatus)
	fmt.Println("\nSorted by Status:")
	PrintTickets(sortedByStatus)
}
