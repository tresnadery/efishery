package domain

import "errors"

var(
	ErrInternalServerError = errors.New("Internal Server Error")
	ErrNotFound = errors.New("Your requested item not found")
	ErrBadParamInput = errors.New("Given Param is not valid")
	ErrUnathorizedToken = errors.New("Unauthorized Access Token")
	ErrCantSignToken = errors.New("Can't signed Token")
	ErrTokenNotFound = errors.New("Token not found")
	ErrUnexpectedSingningMethod = errors.New("Unexpected signing method: %v")	
	ErrConflict = errors.New("Your Item already exist")

)

type ResponseErrorAPI struct{
	Status int `json:"status"`
	Error string `json:"error"`
}

type ResponseError struct{
	Message string `json:"message"`
}