package main

import (
	"fmt"
	"sort"
	"time"
)

const sortModeTime = "time"
const sortModeStatus = "status"

// Changed Ticket struct's field types
type Ticket struct {
	ID          int
	Title       string
	Description string
	Status      string
	CreatedAt   time.Time
}

// implemented builder design pattern to avoid long ticket creation lines and simplify the logic
type TicketBuilder struct {
	nextId int
}

func NewTicketBuilder() *TicketBuilder {
	return &TicketBuilder{
		nextId: 1,
	}
}

func (tb *TicketBuilder) NewTicket(title string, description string, status string) Ticket {
	ticket := Ticket{
		ID:          tb.nextId,
		Title:       title,
		Description: description,
		Status:      status,
		CreatedAt:   time.Now(),
	}
	tb.nextId++
	return ticket
}

/*
 * commented this function out since it's no longer in use but it definetly
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

	tickets := []Ticket{
		ticketBuilder.NewTicket("Ticket 1", "First ticket", "closed"),
		ticketBuilder.NewTicket("Ticket 2", "Second ticket", "open"),
		ticketBuilder.NewTicket("Ticket 3", "Third ticket", "in-progress"),
	}

	sortedByTime := BadSortTickets(tickets, sortModeTime)
	fmt.Println("Sorted by Time:")
	PrintTickets(sortedByTime)

	sortedByStatus := BadSortTickets(tickets, sortModeStatus)
	fmt.Println("\nSorted by Status:")
	PrintTickets(sortedByStatus)

}
