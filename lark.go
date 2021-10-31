package messager

import "github.com/imroc/req"

type LarkMessenger struct {
	url string
}

func NewLarkMessenger(url string) *LarkMessenger {
	return &LarkMessenger{
		url: url,
	}
}

type requestBody struct {
	MsgType string `json:"msg_type"`
	Content struct {
		Text string `json:"text"`
	} `json:"content"`
}

func (m *LarkMessenger) Send(msg string) error {
	body := requestBody{
		MsgType: "text",
		Content: struct {
			Text string `json:"text"`
		}{
			Text: msg,
		},
	}
	resp, err := req.Post(m.url, req.BodyJSON(body))
	return handleResp(resp, err)
}
