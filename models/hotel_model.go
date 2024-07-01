package models

type HotelSearch struct {
	Title       string `json:"title" example:"Hayfield"`
	Name        string `json:"name" example:"Sample Hotel"`
	Description string `json:"description" example:"A sample hotel for testing purposes"`
	Country     string `json:"country" example:"United States"`
	City        string `json:"city" example:"Sample City"`
	State       string `json:"state" example:"Sample State"`
}
type HotelSearchRequest struct {
	HotelSearch
	Offset uint32 `json:"offset" example:"0"`
	Limit  uint32 `json:"limit" example:"10"`
}
