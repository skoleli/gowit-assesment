package main

import (
	"fmt"
	"sort"
	"time"
)

type Ticket struct {
	ID          string
	Title       string
	Description string
	Status      string
	CreatedAt   string
	Timezone    string
}

func parseTimeInUTC(timeStr, timezone string) time.Time {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		loc = time.UTC
	}
	localTime, _ := time.ParseInLocation(time.RFC3339, timeStr, loc)
	return localTime.UTC()
}
func BadSortTickets(tickets []Ticket, mode string) []Ticket {
	if mode == "time" {
		sort.SliceStable(tickets, func(i, j int) bool {
			createdAtI := parseTimeInUTC(tickets[i].CreatedAt, tickets[i].Timezone)
			createdAtJ := parseTimeInUTC(tickets[j].CreatedAt, tickets[j].Timezone)
			return createdAtI.After(createdAtJ)
		})
	} else if mode == "status" {
		sort.SliceStable(tickets, func(i, j int) bool {
			if tickets[i].Status == tickets[j].Status {
				createdAtI := parseTimeInUTC(tickets[i].CreatedAt, tickets[i].Timezone)
				createdAtJ := parseTimeInUTC(tickets[j].CreatedAt, tickets[j].Timezone)
				return createdAtI.After(createdAtJ)
			}
			return tickets[i].Status < tickets[j].Status
		})
	} else {
		// Default sorting by time
		sort.SliceStable(tickets, func(i, j int) bool {
			createdAtI := parseTimeInUTC(tickets[i].CreatedAt, tickets[i].Timezone)
			createdAtJ := parseTimeInUTC(tickets[j].CreatedAt, tickets[j].Timezone)
			return createdAtI.After(createdAtJ)
		})
	}
	return tickets
}
func main() {
	tickets := []Ticket{
		{ID: "1", Title: "Ticket 1", Description: "First ticket", Status: "closed", CreatedAt: "2023-06-05T10:00:00-04:00", Timezone: "America/New_York"},
		{ID: "2", Title: "Ticket 2", Description: "Second ticket", Status: "open", CreatedAt: "2024-06-05T15:00:00+02:00", Timezone: "Europe/Berlin"},
		{ID: "3", Title: "Ticket 3", Description: "Third ticket", Status: "in-progress", CreatedAt: "2024-06-05T12:00:00+09:00", Timezone: "Asia/Tokyo"},
	}
	sortedByTime := BadSortTickets(tickets, "time")
	fmt.Println("Sorted by Time:")
	for _, ticket := range sortedByTime {
		fmt.Printf("%+v\n", ticket)
	}
	sortedByStatus := BadSortTickets(tickets, "status")
	fmt.Println("\nSorted by Status:")
	for _, ticket := range sortedByStatus {
		fmt.Printf("%+v\n", ticket)
	}
}
