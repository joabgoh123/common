package customerrors

type HttpServerError struct {
	statusCode int
	message string
	detail string
}

func (se *HttpServerError) Error() string {
	return se.detail
}
