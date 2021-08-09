package models

type PersonModel struct{}

func (m *PersonModel) GetPersonByID(personID string, age int) (*Person, error) {
	return &Person{
		Id:   personID,
		Name: "João Ribeiro",
		Age:  age,
	}, nil
}
