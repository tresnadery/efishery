package domain

import "errors"

var(
	ErrInternalServerError = errors.New("Internal Server Error")
	ErrNotFound = errors.New("Your requested item not found")
	ErrBadParamInput = errors.New("Given Param is not valid")
	ErrUnathorizedToken = errors.New("Unauthorized Access Token")
	ErrConflict = errors.New("Your Item already exist")
	ErrPasswordIsIncorrect = errors.New("Phone number or password is incorrect")
	ErrCantSignToken = errors.New("Can't signed Token")
	ErrTokenNotFound = errors.New("Token not found")
	ErrUnexpectedSingningMethod = errors.New("Unexpected signing method: %v")
)