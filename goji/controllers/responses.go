package controllers

type PersonResponse struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
}

type AddressResponse struct {
	Id string `json:"id"`
	Country string `json:"country"`
	City string `json:"city"`
	Street string `json:"street"`
	Number int `json:"number"`
}

type ErrorResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
}
