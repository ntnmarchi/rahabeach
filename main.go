package main

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type FlightTickets struct {
	Tickets [][]string `json:"tickets"`
}

func reconstructItinerary(tickets [][]string) ([]string, error) {
	if len(tickets) == 0 {
		return nil, errors.New("no tickets to process")
	}

	startPoints := make(map[string]string)
	endPoints := make(map[string]bool)
	for _, ticket := range tickets {
		start, end := ticket[0], ticket[1]
		startPoints[start] = end
		endPoints[end] = true
	}

	// Determine the starting point: it should be a point that is not an end point
	var start string
	for s := range startPoints {
		if _, found := endPoints[s]; !found {
			start = s
			break
		}
	}

	if start == "" {
		return nil, errors.New("invalid ticket input: no clear starting point")
	}

	// Build the itinerary
	itinerary := []string{}
	current := start
	for {
		itinerary = append(itinerary, current)
		next, exists := startPoints[current]
		if !exists {
			break
		}
		current = next
	}

	// Ensure all tickets are used, indicating a valid path
	if len(itinerary) != len(tickets)+1 {
		return nil, errors.New("disconnected itinerary")
	}

	return itinerary, nil
}

func handlePostItinerary(c echo.Context) error {
	var flightTickets FlightTickets

	if err := c.Bind(&flightTickets); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input format")
	}

	itinerary, err := reconstructItinerary(flightTickets.Tickets)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, itinerary)
}

func main() {
	e := echo.New()                           // Create a new Echo instance
	e.POST("/itinerary", handlePostItinerary) // Bind the POST request to handlePostItinerary handler

	// Start the server on port 8080 and log fatal errors
	e.Logger.Fatal(e.Start(":8080"))
}
