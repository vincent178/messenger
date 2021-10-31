package messager

import (
	"fmt"

	"github.com/imroc/req"
)

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

func handleResp(resp *req.Resp, err error) error {
	if err != nil {
		return err
	}

	if resp.Response().StatusCode >= 200 {
		body, _ := resp.ToString()
		return &resperr{
			code: resp.Response().StatusCode,
			body: body,
		}
	}

	return nil
}
