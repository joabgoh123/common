package customerrors

import (
	"fmt"
	"net/http"
)

type BaseErrorHandler func(http.ResponseWriter, *http.Request) error

func BaseErrorHandlerWrapper(fn func(http.ResponseWriter, *http.Request) error) BaseErrorHandler{
	return BaseErrorHandler(fn)
}

func (fn BaseErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn(w, r)
	if err != nil {
		fmt.Println("Hello world!")
	}
}
