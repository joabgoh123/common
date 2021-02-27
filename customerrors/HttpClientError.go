package customerrors

import (
	"encoding/json"
	"fmt"
)

type HttpClientError struct{
	statusCode int
	message string
	detail string
}

func (ce *HttpClientError) Error() string {
	return ce.detail
}

func (ce *HttpClientError) ResponseBody() ([]byte, error) {
	body, err := json.Marshal(ce)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing response body: %v", err)
	}
	return body, nil
}

func (ce *HttpClientError) ResponseHeaders() (int, map[string]string)  {
	return ce.statusCode, map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	}
}

