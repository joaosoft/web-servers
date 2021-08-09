package models

import "time"

type AddressModel struct{}

func (m *AddressModel) GetPersonAddressByID(personID, addressID string) (*Address, error) {
	// do something
	<-time.After(time.Millisecond * 10)

	return &Address{
		Id:      addressID,
		Country: "Portugal",
		City:    "Porto",
		Street:  "Rua da calçada",
		Number:  7,
	}, nil
}
