package errors

import "net/http"

var DefaultStatusCodeMap = map[string]int{
	RequestTimeoutError:	http.StatusRequestTimeout,
	BadRequestError:		http.StatusBadRequest,
	ValidationError:		http.StatusUnprocessableEntity,
	InternalServerError:	http.StatusInternalServerError,
	UnauthorizedError:		http.StatusUnauthorized,
	NotFoundError:			http.StatusNotFound,
	ForbiddenError:			http.StatusForbidden,

	JsonUnmarshalError:		http.StatusBadRequest,
	JsonMarshalError:		http.StatusBadRequest,

	ProtobufUnmarshalError:	http.StatusBadRequest,
	ProtobufMarshalError:	http.StatusBadRequest,

	PostgresNotFoundError:	http.StatusNotFound,
	PostgresInternalError:	http.StatusInternalServerError,
	PostgresAlreadyExists:	http.StatusConflict,

	RedisNotFoundError:		http.StatusNotFound,
	RedisInternalError:		http.StatusInternalServerError,
}
