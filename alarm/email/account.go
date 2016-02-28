package email

type Account struct {
	Username string
	Nickname string
	Password string
	Host     string
}

func NewAccount(user, nick, passwd, host string) *Account {
	a := &Account{
		Username: user,
		Nickname: nick,
		Password: passwd,
		Host:     host,
	}
	return a
}
