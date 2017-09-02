package sms

import (
	"bytes"
	"encoding/json"
	"net/url"

	"github.com/ckeyer/commons/httpclient"
	"github.com/ckeyer/commons/util"
)

const (
	AliAPIServer = "https://sms.aliyuncs.com/"
)

type AliSmsClient struct {
	cli *httpclient.Client

	AccessKeyId     string
	AccessKeySecret string
}

type AliReqBody struct {
	AccessKeyId  string `json:"AccessKeyId" form:"AccessKeyId"`
	Action       string `json:"Action" form:"Action"`
	SignName     string `json:"SignName" form:"SignName"`
	TemplateCode string `json:"TemplateCode" form:"TemplateCode"`
	RecNum       string `json:"RecNum" form:"RecNum"`
	ParamString  string `json:"ParamString" form:"ParamString"`

	Format           string `json:"Format" form:"Format"`
	Version          string `json:"Version" form:"Version"`
	Signature        string `json:"Signature" form:"Signature"`
	SignatureMethod  string `json:"SignatureMethod" form:"SignatureMethod"`
	SignatureNonce   string `json:"SignatureNonce" form:"SignatureNonce"`
	SignatureVersion string `json:"SignatureVersion" form:"SignatureVersion"`
	Timestamp        string `json:"Timestamp" form:"Timestamp"`
}

func NewAliSmsCli(accessKeyId, accessKeySecret string) *AliSmsClient {
	return &AliSmsClient{
		cli: httpclient.NewClient(),

		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
	}
}

func (a *AliSmsClient) newAliReqBody() *AliReqBody {
	body := &AliReqBody{
		AccessKeyId:      a.AccessKeyId,
		Action:           "SingleSendSms",
		Format:           "JSON",
		Version:          "2016-09-27",
		Signature:        "Pc5WB8gokVn0xfeu%2FZV%2BiNM1dgI%3D",
		SignatureMethod:  "HMAC-SHA1",
		SignatureVersion: "1.0",
		Timestamp:        "2015-11-23T12:00:00Z",
	}
}

func (a *AliSmsClient) baseReq() *url.URL {
	u, _ := url.Parse(AliAPIServer)
	q := u.Query()

	q.Set("Format", a.Format)
	q.Set("Version", a.Version)
	q.Set("Signature", a.Signature)
	q.Set("SignatureMethod", a.SignatureMethod)
	q.Set("SignatureNonce", util.RandomUUID())
	q.Set("SignatureVersion", a.SignatureVersion)
	q.Set("AccessKeyId", a.AccessKeyId)
	q.Set("Timestamp", a.Timestamp)

	u.RawQuery = q.Encode()
	return u
}

func (a *AliSmsClient) SendSms(tmplID string, data interface{}) error {
	buf := &bytes.Buffer{}
	err := json.NewEncoder(buf).Encode(data)
	if err != nil {
		return err
	}

	a.constructReqBody(tmplID, buf.String())
}
