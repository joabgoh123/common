package customerrors

type HttpClientError struct {
	statusCode int
	message    string
	detail     string
}

func (ce *HttpClientError) Error() string {
	return ce.detail
}

func (ce *HttpClientError) ResponseBody() map[string]string {
	return map[string]string{"error": ce.message, "detail": ce.detail}
}

func (ce *HttpClientError) ResponseHeaders() (int, map[string]string) {
	return ce.statusCode, map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	}
}
