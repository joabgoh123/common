package customerrors

type ServerError interface {
	Error() string
	ErrorLog()
	ResponseHeaders() (int, map[string]string)
	ResponseBody() map[string]string
}
