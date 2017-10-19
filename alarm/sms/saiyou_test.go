package sms

import "testing"

type SaiyouRemind struct {
	Name    string `json:"name"`
	Task    string `json:"task"`
	Content string `json:"time"`
}

type SaiyouNoWater struct {
	Content string `json:"time"`
}

const (
	sms_to        = ""
	saiyou_prj    = ""
	saiyou_appid  = ""
	saiyou_appkey = ""
)

func TestSaiyouSend(t *testing.T) {
	cli := NewSaiyouSMS(saiyou_appid, saiyou_appkey)
	err := cli.Send(sms_to, saiyou_prj, SaiyouRemind{"王大猛子", "打酱油", "愿世界充满爱"})
	if err != nil {
		t.Error(err)
		return
	}
}
