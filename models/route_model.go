package models

// Route document model
type Schedule struct {
	Day    int    `json:"day" example:"1"`
	Flight string `json:"flight" example:"XYZ123"`
	UTC    string `json:"utc" example:"14:30"`
}

type Route struct {
	Airline            string     `json:"airline,omitempty" binding:"required" example:"AF"`
	Airline_id         string     `json:"airlineid,omitempty" binding:"required" example:"airline_10"`
	SourceAirport      string     `json:"sourceairport,omitempty" binding:"required" example:"SFO"`
	DestinationAirport string     `json:"destinationairport,omitempty" binding:"required" example:"JFK"`
	Stops              int        `json:"stops,omitempty" example:"0"`
	Equipment          string     `json:"equipment,omitempty" example:"CRJ"`
	Schedule           []Schedule `json:"schedule,omitempty"`
	Distance           float64    `json:"distance,omitempty" example:"4151.79"`
}