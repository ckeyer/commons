package useragent

import (
	"github.com/Masterminds/semver"
	"github.com/ckeyer/commons/crypto"
)

func RandUserAgent() string {
	return Commons[crypto.RandomInt(0, len(Commons)-1)]
}

type UserAgent struct {
	OS             string
	OSVersion      semver.Version
	CPUType        string
	Browser        string
	BrowserVersion semver.Version
	Language       string
	Plugins        string
}

func NewFirefox() *UserAgent {
	return nil
}

func (u *UserAgent) String() string {
	return ""
}
