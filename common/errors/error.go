package errors

import (
	"fmt"
	"net/http"
)

const (
	RequestTimeoutError		string = "RequestTimeoutError"
	BadRequestError			string = "BadRequestError"
	ValidationError			string = "ValidationError"
	InternalServerError		string = "InternalServerError"
	NotFoundError			string = "NotFoundError"
	UnauthorizedError		string = "UnauthorizedError"
	ForbiddenError			string = "ForbiddenError"

	JsonUnmarshalError		string = "JsonUnmarshalError"
	JsonMarshalError		string = "JsonMarshalError"

	ProtobufUnmarshalError	string = "ProtobufUnmarshalError"
	ProtobufMarshalError	string = "ProtobufMarshalError"

	PostgresNotFoundError	string = "PostgresNotFoundError"
	PostgresInternalError	string = "PostgresInternalError"
	PostgresAlreadyExists	string = "PostgresAlreadyExists"

	RedisNotFoundError		string = "RedisNotFoundError"
	RedisInternalError		string = "RedisInternalError"
)

type Error struct {
	Message          string                 `json:"message"`
	Type             string                 `json:"type"`
	ValidationErrors map[string]interface{} `json:"errors"`
}

func New(message, errorType string) Error {
	return Error{
		Message:          message,
		Type:             errorType,
		ValidationErrors: make(map[string]interface{}),
	}
}

func (e Error) StatusCodeFromMap() int32 {
	code, ok := DefaultStatusCodeMap[e.Type]
	if !ok {
		code = http.StatusInternalServerError
	}

	return int32(code)
}

func (e Error) WithValidationError(name, value, location, message string) Error {
	e.ValidationErrors[name] = map[string]string{
		"location": location,
		"value":    value,
		"message":  message,
	}

	return e
}

func (e Error) IsValidation() bool {
	return e.Type == ValidationError
}

func (e Error) Error() string {
	return fmt.Sprintf("status: %d, message: %s, type: %s", e.StatusCodeFromMap(), e.Message, e.Type)
}
