package errs

import (
	"net/http"
)

var (
	ErrInvalidCredentials         = NewAPIError(http.StatusUnauthorized, "invalid credentials")
	ErrInvalidToken               = NewAPIError(http.StatusUnauthorized, "invalid token")
	ErrEmailAlreadyTaken          = NewAPIError(http.StatusUnprocessableEntity, "email has been already taken")
	ErrEntityDoesNotBelongsToUser = NewAPIError(http.StatusForbidden, "entity doesn't belong to given user")
	ErrUnableToBindRequest        = NewAPIError(http.StatusBadRequest, "unable to bind given request to struct")
)
