package dto

type Response[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func NewResponseSuccess[T any](message string, data T) Response[T] {
	return Response[T]{
		Message: message,
		Data:    data,
	}
}

func NewResponseError(message string) Response[string] {
	return Response[string]{
		Message: message,
		Data:    "",
	}
}

func NewResponseErrorWithData[T any](message string, data T) Response[T] {
	return Response[T]{
		Message: message,
		Data:    data,
	}
}
