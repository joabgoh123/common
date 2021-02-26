package customerrors

import "net/http"

type BaseErrorHandler func(http.ResponseWriter, *http.Request) error

func (fn BaseErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn(w, r)
	if err != nil {

	}
}