package responses

import "src/models"

type TravelSampleResponse struct {
	Status         int         `json:"status"`
	Message        string      `json:"message"`
	CollectionData interface{} `json:"data"`
}

type TravelSampleResponseForAirline struct {
	Status         int                          `json:"status" example:"200"`
	Message        string                       `json:"message" example:"Success"`
	CollectionData models.RequestBodyForAirline `json:"data"`
}
