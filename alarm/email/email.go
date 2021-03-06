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

	A_LEVEL_DEBUG   = "debug"
	A_LEVEL_WARM    = "warm"
	A_LEVEL_ERROR   = "error"
	A_LEVEL_SUCCESS = "success"
	A_LEVEL_INFO    = "info"
)

type Email struct {
	acc  *Account
	tmpl map[string]interface{}
}

func SendMail(account Account, title, body string, send_to ...string) error {
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
	err := smtp.SendMail(account.Host, account.Auth, account.Username, send_to, content)
	return err
}
