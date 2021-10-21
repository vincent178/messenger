package messager

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/imroc/req"
)

// https://developers.dingtalk.com/document/robots/custom-robot-access
type DingdingMessenger struct {
	token string
	sec   string
}

func NewDingdingMessager(token, sec string) *DingdingMessenger {
	return &DingdingMessenger{
		token: token,
		sec:   sec,
	}
}

type DingdingResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (m *DingdingMessenger) Send(msg string) error {
	param := req.Param{
		"msgtype": "text",
		"text":    req.Param{"content": msg},
	}

	ts := time.Now().Unix() * 1000
	strToEncode := strconv.Itoa(int(ts)) + "\n" + m.sec
	h := hmac.New(sha256.New, []byte(m.sec))
	h.Write([]byte(strToEncode))
	sign := base64.StdEncoding.EncodeToString(h.Sum(nil))

	q := url.Values{}
	q.Add("access_token", m.token)
	q.Add("timestamp", strconv.Itoa(int(ts)))
	q.Add("sign", sign)

	resp, err := req.Post("https://oapi.dingtalk.com/robot/send?"+q.Encode(), req.BodyJSON(param))
	if err != nil {
		return err
	}

	if resp.Response().StatusCode != 200 {
		body, _ := resp.ToString()
		return &resperr{
			code: resp.Response().StatusCode,
			body: body,
		}
	}

	r := DingdingResponse{}
	resp.ToJSON(&r)

	if r.ErrCode != 0 {
		return &resperr{
			code: resp.Response().StatusCode,
			body: fmt.Sprintf("errcode: %d errmsg: %s", r.ErrCode, r.ErrMsg),
		}
	}

	return nil
}
