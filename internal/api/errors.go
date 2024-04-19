package api

type QueryParamRequired struct {
	msg string
}

func (v QueryParamRequired) Error() string {
	return v.msg
}
