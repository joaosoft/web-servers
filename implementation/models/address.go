package models

type AddressModel struct{}

func (m *AddressModel) GetPersonAddressByID(personID, addressID string) (*Address, error) {
	return &Address{
		Id:      addressID,
		Country: "Portugal",
		City:    "Porto",
		Street:  "Rua da calçada",
		Number:  7,
	}, nil
}
