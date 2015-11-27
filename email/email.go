package email

import (
	"errors"
	"net/smtp"
	"strings"
)

var (
	ERR_NO_SEND_LIST = errors.New("no found send list")
)

const (
	CONTENT_TYPE_HTML = "Content-Type: text/html; charset=UTF-8"
	CONTENT_TYPE_TEXT = "Content-Type: text/plain; charset=UTF-8"
)

func SendMail(account Account, title, body string, send_to ...string) error {
	hp := strings.Split(account.Host, ":")
	auth := smtp.PlainAuth("", account.Username, account.Password, hp[0])
	content_type := CONTENT_TYPE_TEXT
	if strings.HasPrefix(body, "<!DOCTYPE html>") || strings.HasPrefix(body, "<html") {
		content_type = CONTENT_TYPE_HTML
	}
	if len(send_to) == 0 {
		return ERR_NO_SEND_LIST
	}
	to := strings.Join(send_to, ";")
	content := []byte("To: " + to + "\r\nFrom: " + account.Nickname + "<" + account.Username + ">\r\nSubject: " + title + "\r\n" +
		content_type + "\r\n\r\n" + body)
	err := smtp.SendMail(account.Host, auth, account.Username, send_to, content)
	return err
}
