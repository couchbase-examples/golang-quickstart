package responses

type ProfileResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Profile interface{} 			`json:"data"`
}
