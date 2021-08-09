package models

import "net/http"

type ErrorModel struct{}

func (m *ErrorModel) GetErrorByID(errorID int) (*Error, error) {
	return &Error{
		Code:    errorID,
		Message: http.StatusText(errorID),
	}, nil
}
