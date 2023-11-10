package models

// Airline document model
type Airline struct {
	Callsign string `json:"callsign,omitempty" example:"SAF"`
	Country  string `json:"country,omitempty" binding:"required" example:"United States"`
	IATA     string `json:"iata,omitempty" example:"SA"`
	ICAO     string `json:"icao,omitempty" binding:"required" example:"SAF"`
	Name     string `json:"name,omitempty" binding:"required" example:"SampleName"`
}

// Airport document model
type Geo struct {
	Alt float64 `json:"alt" example:"48.864716"`
	Lat float64 `json:"lat" example:"2.349014"`
	Lon float64 `json:"lon" example:"92.0"`
}

type Airport struct {
	AirportName string `json:"airportname,omitempty" binding:"required" example:"SampleAirport"`
	City        string `json:"city,omitempty" binding:"required" example:"SampleCity"`
	Country     string `json:"country,omitempty" binding:"required" example:"United Kingdom"`
	FAA         string `json:"faa,omitempty" binding:"required" example:"SAA"`
	GEO         Geo    `json:"geo,omitempty"`
	ICAO        string `json:"icao,omitempty" example:"SAAA"`
	TZ          string `json:"tz,omitempty" example:"Europe/Paris"`
}

// Route document model
type Schedule struct {
	Day    int    `json:"day" example:"1"`
	Flight string `json:"flight" example:"XYZ123"`
	UTC    string `json:"utc" example:"14:30"`
}

type Route struct {
	Airline            string     `json:"airline,omitempty" binding:"required" example:"AF"`
	Airline_id         string     `json:"airline_id,omitempty" binding:"required" example:"airline_10"`
	SourceAirport      string     `json:"sourceairport,omitempty" binding:"required" example:"SFO"`
	DestinationAirport string     `json:"destinationairport,omitempty" binding:"required" example:"JFK"`
	Stops              int        `json:"stops,omitempty" example:"0"`
	Equipment          string     `json:"equipment,omitempty" example:"CRJ"`
	Schedule           []Schedule `json:"schedule,omitempty"`
	Distance           float64    `json:"distance,omitempty" example:"4151.79"`
}

type Destination struct {
	DestinationAirport string `json:"destinationairport" example:"JFK"`
}
