package sms

import (
	"testing"
)

var (
	alikeyid     = "LTAIVsVc3tUWjWOU"
	alikeysecret = "WDm4dQK4i77vNRdrkTu3eZgBF0w18P"
)

func TestAliSms(t *testing.T) {
	cli := NewAliSmsCli(alikeyid, alikeysecret)
}
