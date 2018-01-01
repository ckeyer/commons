package sms

import (
	"net/url"

	"github.com/ckeyer/commons/crypto/uuid"

	"github.com/ckeyer/commons/httpcli"
)

const (
	AliAPIServer = "https://sms.aliyuncs.com/"
)

type AliSmsClient struct {
	cli *httpcli.Client

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
		cli: httpcli.NewClient(),

		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
	}
}

func (a *AliSmsClient) newAliReqBody() *AliReqBody {
	return &AliReqBody{
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

func (a *AliSmsClient) baseReq(req *AliReqBody) *url.URL {
	u, _ := url.Parse(AliAPIServer)
	q := u.Query()

	q.Set("Format", req.Format)
	q.Set("Version", req.Version)
	q.Set("Signature", req.Signature)
	q.Set("SignatureMethod", req.SignatureMethod)
	q.Set("SignatureNonce", uuid.NewV4().String())
	q.Set("SignatureVersion", req.SignatureVersion)
	q.Set("AccessKeyId", a.AccessKeyId)
	q.Set("Timestamp", req.Timestamp)

	u.RawQuery = q.Encode()
	return u
}

func (a *AliSmsClient) SendSms(tmplID string, data interface{}) error {
	// buf := &bytes.Buffer{}
	// err := json.NewEncoder(buf).Encode(data)
	// if err != nil {
	// 	return err
	// }

	return nil
}
