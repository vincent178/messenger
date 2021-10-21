package messager

import "fmt"

type Messenger interface {
	Send(msg string) error
}

type resperr struct {
	code int
	body string
}

func (r *resperr) Error() string {
	return fmt.Sprintf("http error code %d body %s", r.code, r.body)
}
