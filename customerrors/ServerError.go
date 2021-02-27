package customerrors

type ServerError interface {
	Error() string
}
