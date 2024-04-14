package constant

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

var ERROR_MESSAGE = map[int]error{
	fiber.StatusOK: errors.New("success"),

	// Client errors (4xx)
	fiber.StatusBadRequest:           errors.New("bad request, please check your request parameters"),
	fiber.StatusUnauthorized:         errors.New("unauthorized, please provide valid credentials"),
	fiber.StatusForbidden:            errors.New("forbidden, you don't have permission to access this resource"),
	fiber.StatusNotFound:             errors.New("not found, the requested resource does not exist"),
	fiber.StatusMethodNotAllowed:     errors.New("method not allowed, please use a valid HTTP method"),
	fiber.StatusRequestTimeout:       errors.New("request timeout, please try again later"),
	fiber.StatusConflict:             errors.New("conflict, the request conflicts with the current state of the server"),
	fiber.StatusPreconditionFailed:   errors.New("precondition failed, one or more request headers are invalid"),
	fiber.StatusUnsupportedMediaType: errors.New("unsupported media type, please use a supported content type"),

	// Server errors (5xx)
	fiber.StatusInternalServerError: errors.New("internal server error, please try again later"),
	fiber.StatusNotImplemented:      errors.New("not implemented, the server does not support the requested functionality"),
	fiber.StatusServiceUnavailable:  errors.New("service unavailable, the server is currently unable to handle the request"),
}
