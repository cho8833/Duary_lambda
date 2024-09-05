package util

type ServerResponse[T any] struct {
	message *string
	status  int64
	data    *T
	error   bool
}

func ErrorResponse(message string, statusCode int64) ServerResponse[any] {
	return ServerResponse[any]{
		&message,
		statusCode,
		nil,
		true,
	}
}

func ResponseFromError(error error, statusCode int64) ServerResponse[any] {
	errorMessage := error.Error()
	return ServerResponse[any]{
		&errorMessage,
		statusCode,
		nil,
		true,
	}
}

func ResponseWithData[T any](data T) ServerResponse[T] {
	okString := "OK"
	return ServerResponse[T]{
		&okString,
		200,
		&data,
		false,
	}
}

func SUCCESS() ServerResponse[any] {
	okString := "OK"

	return ServerResponse[any]{
		&okString,
		200,
		nil,
		false,
	}
}
