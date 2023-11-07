package responses

import "src/models"

type TravelSampleResponseForAirport struct {
	Status         int                          `json:"status"`
	Message        string                       `json:"message"`
	CollectionData models.RequestBodyForAirline `json:"data"`
}

func (r *TravelSampleResponseForAirline) SetStatus(status int) {
	r.Status = status
}

func (r *TravelSampleResponseForAirline) SetMessage(message string) {
	r.Message = message
}

func (r *TravelSampleResponseForAirline) SetCollectionData(data models.RequestBodyForAirline) {
	r.CollectionData = data
}

func (r *TravelSampleResponseForAirline) GetResponse() interface{} {
	return r
}