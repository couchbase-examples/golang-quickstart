package models

type RequestBody struct {
	Pid string `json:"Pid,omitempty"`
	FirstName string `json:"FirstName,omitempty"`
	LastName string `json:"LastName,omitempty"`
	Email string `json:"Email" binding:"required"`
	Password string `json:"password" binding:"required"`

}