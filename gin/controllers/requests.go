package controllers

type GetPersonByIDRequest struct {
	IdPerson string `json:"id_person"`
	Age int `json:"age"`
}

type GetPersonAddressByIDRequest struct {
	IdPerson string `json:"id_person"`
	IdAddress string `json:"id_address"`
}
