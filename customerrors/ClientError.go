package customerrors

type ClientError interface {
	Error() string
	ResponseBody() map[string]string
	ResponseHeaders() (int, map[string]string)
}
