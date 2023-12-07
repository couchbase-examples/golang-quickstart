package models

// Airline document model
type Airline struct {
	Callsign string `json:"callsign,omitempty" example:"SAF"`
	Country  string `json:"country,omitempty" binding:"required" example:"United States"`
	IATA     string `json:"iata,omitempty" example:"SA"`
	ICAO     string `json:"icao,omitempty" binding:"required" example:"SAF"`
	Name     string `json:"name,omitempty" binding:"required" example:"SampleName"`
}
