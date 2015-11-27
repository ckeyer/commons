package email

import (
	"fmt"
	"testing"
)

func TestSendEmail(t *testing.T) {
	return
	err := SendMail(defaultAccount, "邮件发送测试", "this is a test email", "dev@ckeyer.com", "wangcj1214@gmail.com")
	if err != nil {
		t.Error(err)
	}
}

func TestSetContent(t *testing.T) {
	bs := SetContent("hello world")
	fmt.Println(string(bs))
	fmt.Println("over")
	t.Error("...")
}
