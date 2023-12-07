package models

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