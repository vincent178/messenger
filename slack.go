package messenger

import "github.com/imroc/req"

type SlackMessenger struct {
	url string
}

func NewSlackMessenger(url string) *SlackMessenger {
	return &SlackMessenger{
		url: url,
	}
}

func (m *SlackMessenger) Send(msg string) error {
	resp, err := req.Post(m.url, req.BodyJSON(map[string]string{"text": msg}))
	return handleResp(resp, err)
}
