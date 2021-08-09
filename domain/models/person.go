package models

import "time"

type PersonModel struct{}

func (m *PersonModel) GetPersonByID(personID string, age int) (*Person, error) {
	// do something
	<-time.After(time.Millisecond * 10)

	return &Person{
		Id:   personID,
		Name: "João Ribeiro",
		Age:  age,
	}, nil
}
