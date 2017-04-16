package email

import (
	"fmt"
	"testing"
)

func TestSendEmail(t *testing.T) {
	a := defaultAccount
	err := SendMail(*a, "邮件发送测试", "this is a test email", "me@ckeyer.com", "chuanjian@staff.sina.com.cn")
	if err != nil {
		t.Error(err.Error())
	}

}

func TestSetContent(t *testing.T) {
	return
	bs := SetContent("hello world")
	fmt.Println(string(bs))
	fmt.Println("over")
	t.Error("...")
}
