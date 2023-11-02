package models


// Airline document model
type RequestBodyForAirline struct {
	Callsign  *string `json:"callsign,omitempty"`
	Country   string  `json:"country" binding:"required"`
	IATA      *string `json:"iata,omitempty"`
	ICAO      string  `json:"icao" binding:"required"`
	Name      string  `json:"name" binding:"required"`
}

// Airport document model
type RequestBodyForAirport struct {
    AirportName string `json:"airportname"`
    City        string `json:"city"`
    Country     string `json:"country"`
    FAA         *string `json:"faa,omitempty"`
    GEO         struct {
        Alt float64 `json:"alt"`
        Lat float64 `json:"lat"`
        Lon float64 `json:"lon"`
    } `json:"geo"`
    ICAO        *string `json:"icao,omitempty"`
    TZ          string `json:"tz"`
}

// Route document model
type RequestBodyForRoute struct {
    Airline            string `json:"airline"`
    AirlineID          string `json:"airlineid"`
    DestinationAirport string `json:"destinationairport"`
    Distance           float64 `json:"distance"`
    Equipment          string `json:"equipment"`
    Schedule           []struct {
        Day    int    `json:"day"`
        Flight string `json:"flight"`
        UTC    string `json:"utc"`
    } `json:"schedule"`
    SourceAirport string `json:"sourceairport"`
    Stops         int    `json:"stops"`
}