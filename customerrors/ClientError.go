package customerrors

type ClientError struct{
	statusCode int
	msg string
	detail string
}

func(ce *ClientError) Error() string{
	return ce.msg
}

func NewClientError(statusCode int, msg, detail string) *ClientError {
	return &ClientError{statusCode, msg, detail}
}


