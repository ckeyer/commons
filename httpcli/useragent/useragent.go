package useragent

import (
	"math/rand"
	"time"
)

const (
	UserAgentHeader = "User-Agent"
)

var (
	allCommons = []string{}
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func init() {
	if len(allCommons) > 0 {
		return
	}
	for _, v := range [][]string{
		Windows,
		Linuxs,
		IPads,
		IPhones,
		Macintosh,
		Andorids,
	} {
		allCommons = append(allCommons, v...)
	}
}

func Commons() {

}

func RandUserAgent() string {
	return allCommons[rand.Intn(len(allCommons))]
}
