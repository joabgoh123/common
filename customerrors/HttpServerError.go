package customerrors

import (
	log "github.com/sirupsen/logrus"
  )

type HttpServerError struct {
	statusCode int
	errorID    int
	message    string
	detail     string
}

func (se *HttpServerError) Error() string {
	return se.detail
}

func (se *HttpServerError) ErrorLog() {
	log.WithFields(log.Fields{
		"errors": &se,
	})
}

func (se *HttpServerError) ResponseHeaders() (int, map[string]string){
	return se.statusCode, map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	}
}

func (se *HttpServerError) ResponseBody() map[string]string {
	 return map[string]string{"error": se.message}
}
