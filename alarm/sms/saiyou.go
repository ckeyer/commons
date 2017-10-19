package sms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type SaiyouSMS struct {
	apiurl string
	cli    *http.Client

	AppID  string
	AppKey string
}

type SendResponse struct {
	Status     string `json:"status"`
	SendId     string `json:"send_id"`
	Fee        int    `json:"fee"`
	SMSCredits string `json:"sms_credits"`
	Message    string `json:"msg"`
}

func NewSaiyouSMS(appid, appkey string) *SaiyouSMS {
	return &SaiyouSMS{
		apiurl: "https://api.mysubmail.com/message/xsend.json",
		cli:    http.DefaultClient,
		AppID:  appid,
		AppKey: appkey,
	}
}

func (m *SaiyouSMS) Send(to, tpl string, data interface{}) error {
	buf := &bytes.Buffer{}
	err := json.NewEncoder(buf).Encode(data)
	if err != nil {
		return err
	}
	q := url.Values{}
	q.Set("appid", m.AppID)
	q.Set("signature", m.AppKey)
	q.Set("to", to)
	q.Set("project", tpl)
	q.Set("vars", buf.String())

	resp, err := m.cli.PostForm(m.apiurl, q)
	if err != nil {
		return err
	}
	ret := &SendResponse{}
	err = json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		return err
	}

	if ret.Status != "success" {
		return fmt.Errorf("saiyou send failed, %s", ret.Message)
	}
	return nil
}
