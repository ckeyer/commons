package ssh

import (
	"golang.org/x/crypto/ssh"
)

func runtest() {
	ssh.Dial("tcp", addr, config)
	ssh.NewClient(c, chans, reqs)
}
