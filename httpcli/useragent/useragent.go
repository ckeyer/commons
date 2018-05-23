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
	pcCommons  = []string{}
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
		WinPhones,
		Macintosh,
		Andorids,
	} {
		allCommons = append(allCommons, v...)
	}

	for _, v := range [][]string{
		Windows,
		Linuxs,
		Macintosh,
	} {
		pcCommons = append(pcCommons, v...)
	}
}

func Commons() {

}

func RandUserAgent() string {
	return allCommons[rand.Intn(len(allCommons))]
}

func PCUserAgent() string {
	return pcCommons[rand.Intn(len(pcCommons))]
}
