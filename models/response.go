package models 

type Response struct {
	Status  string `json:"status"`
	Data    any    `json:"data,omitempty"` 
}