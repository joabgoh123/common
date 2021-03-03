package customerrors

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

type HttpServerError struct {
	StatusCode int
	ErrorID    int
	Message    string
	Detail     string
}

func (se *HttpServerError) Error() string {
	return se.Detail
}

func (se *HttpServerError) ErrorLog() {
	log.WithFields(log.Fields{
		"errors": se,
	})
}

func (se *HttpServerError) ResponseHeaders() (int, map[string]string){
	return se.StatusCode, map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	}
}

func (se *HttpServerError) ResponseBody() map[string]string {
	 return map[string]string{"error": se.Message}
}
