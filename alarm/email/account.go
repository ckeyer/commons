package email

import (
	"errors"
	"net/smtp"
	"strings"
)

type Account struct {
	Username string
	Nickname string
	Password string
	Host     string

	Auth smtp.Auth
}

func NewAccount(user, nick, passwd, host string, auth ...smtp.Auth) *Account {
	a := &Account{
		Username: user,
		Nickname: nick,
		Password: passwd,
		Host:     host,
	}
	if len(auth) == 1 {
		a.Auth = auth[0]
	} else {
		a.Auth = a.PlainAuth()
	}
	return a
}

func (a *Account) PlainAuth() smtp.Auth {
	hp := strings.SplitN(a.Host, ":", 2)
	return smtp.PlainAuth("", a.Username, a.Password, hp[0])
}

func (a *Account) CRAMMD5Auth(secret string) smtp.Auth {
	return smtp.CRAMMD5Auth(a.Username, secret)
}

func (a *Account) LoginAuth() smtp.Auth {
	return &loginAuth{a.Username, a.Password}
}

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

// Used for AUTH LOGIN. (Maybe password should be encrypted)
func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch strings.ToLower(string(fromServer)) {
		case "username:":
			return []byte(a.username), nil
		case "password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("unexpected server challenge")
		}
	}
	return nil, nil
}
