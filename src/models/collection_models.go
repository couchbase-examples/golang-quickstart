package models

// Airline document model
type RequestBodyForAirline struct {
	Callsign string `json:"callsign,omitempty" default:"SampleCallsign"`
	Country  string `json:"country" binding:"required" default:"SampleCountry"`
	IATA     string `json:"iata,omitempty" default:"SMP"`
	ICAO     string `json:"icao" binding:"required" default:"SMPL"`
	Name     string `json:"name" binding:"required" default:"SampleName"`
}

// Airport document model
// type RequestBodyForAirport struct {
// 	AirportName string `json:"airportname" binding:"required"`
// 	City        string `json:"city" binding:"required"`
// 	Country     string `json:"country" binding:"required"`
// 	FAA         string `json:"faa,omitempty" binding:"required"`
// 	GEO         struct {
// 		Alt float64 `json:"alt"`
// 		Lat float64 `json:"lat"`
// 		Lon float64 `json:"lon"`
// 	} `json:"geo"`
// 	ICAO string `json:"icao,omitempty"`
// 	TZ   string `json:"tz"`
// }

// Airport document model
type RequestBodyForAirport struct {
	AirportName string `json:"airportname" binding:"required" example:"SampleAirport"`
	City        string `json:"city" binding:"required" example:"SampleCity"`
	Country     string `json:"country" binding:"required" example:"United Kingdom"`
	FAA         string `json:"faa,omitempty" binding:"required" example:"SAA"`
	GEO         struct {
		Alt float64 `json:"alt" example:48.864716`
		Lat float64 `json:"lat" example:2.349014`
		Lon float64 `json:"lon" example:92.0`
	} `json:"geo"`
	ICAO string `json:"icao,omitempty" example:"SAAA"`
	TZ   string `json:"tz" example:"Europe/Paris"`
}

// Route document model
type RequestBodyForRoute struct {
	Airline            string  `json:"airline"`
	AirlineID          string  `json:"airlineid"`
	DestinationAirport string  `json:"destinationairport"`
	Distance           float64 `json:"distance"`
	Equipment          string  `json:"equipment"`
	Schedule           []struct {
		Day    int    `json:"day"`
		Flight string `json:"flight"`
		UTC    string `json:"utc"`
	} `json:"schedule"`
	SourceAirport string `json:"sourceairport"`
	Stops         int    `json:"stops"`
}
