package customerrors

type HttpClientError struct {
	StatusCode int
	Message    string
	Detail     string
}

func (ce *HttpClientError) Error() string {
	return ce.Detail
}

func (ce *HttpClientError) ResponseBody() map[string]string {
	return map[string]string{"error": ce.Message, "Detail": ce.Detail}
}

func (ce *HttpClientError) ResponseHeaders() (int, map[string]string) {
	return ce.StatusCode, map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	}
}
