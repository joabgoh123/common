package customerrors

import (
	"encoding/json"
	"net/http"
)

type BaseErrorHandler func(http.ResponseWriter, *http.Request) error

func (fn BaseErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn(w, r)
	if err != nil {
		switch err := err.(type) {
			
		case ServerError:
			err.ErrorLog()
			statusCode, headers := err.ResponseHeaders()
			body := err.ResponseBody()
			w.Header().Set("Content-Type", headers["Content-Type"])
			w.WriteHeader(statusCode)
			json.NewEncoder(w).Encode(body)

		case ClientError:
			statusCode, headers := err.ResponseHeaders()
			body := err.ResponseBody()
			w.Header().Set("Content-Type", headers["Content-Type"])
			w.WriteHeader(statusCode)
			json.NewEncoder(w).Encode(body)
	}
}
}