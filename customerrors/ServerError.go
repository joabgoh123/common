package customerrors

type ServerError struct{
	statusCode int
	msg string
	detail string
}

func(ce *ServerError) Error() string{
	return ce.msg
}

func NewServerError(statusCode int, msg, detail string) *ServerError {
	return &ServerError{statusCode, msg, detail}
}


