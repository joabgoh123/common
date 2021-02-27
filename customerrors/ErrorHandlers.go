package customerrors

import (
	"fmt"
	"net/http"
)

type BaseErrorHandler func(http.ResponseWriter, *http.Request) error

func (fn BaseErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn(w, r)
	if err != nil {
		fmt.Println("Hello world!")
	}
}