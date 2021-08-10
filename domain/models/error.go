package models

import (
	"net/http"
	"time"
)

type ErrorModel struct{}

func (m *ErrorModel) GetErrorByID(errorID int) (*Error, error) {
	// do something
	<-time.After(time.Millisecond * 10)

	return &Error{
		Code:    errorID,
		Message: http.StatusText(errorID),
	}, nil
}
