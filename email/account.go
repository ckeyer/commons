package email

type Account struct {
	Username string
	Nickname string
	Password string
	Host     string
}

var (
	DefaultAccount = Account{
		Username: "ckeyer_alarm@163.com",
		Nickname: "Ckeyer-Alarm",
		Password: "fpwjxqryyrfqccjr",
		Host:     "smtp.163.com:25",
	}
	TmpAccount = Account{
		Username: "sf_monitor@163.com",
		Nickname: "SF-Monitor",
		Password: "cqwsmdgupjmareyc",
		Host:     "smtp.163.com:25",
	}
	FunxData = Account{
		Username: "funxdata@163.com",
		Nickname: "Funx-Data",
		Password: "orrmtgbgmdxmcsru",
		Host:     "smtp.163.com:25",
	}
)

func NewAccount(user, nick, passwd, host string) *Account {
	a := &Account{
		Username: user,
		Nickname: nick,
		Password: passwd,
		Host:     host,
	}
	return a
}
