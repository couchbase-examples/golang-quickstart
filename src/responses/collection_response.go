package responses

type TravelSampleResponse struct {
	Status         int         `json:"status"`
	Message        string      `json:"message"`
	CollectionData interface{} `json:"data"`
}
