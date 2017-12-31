package redis

import (
	"errors"
	"fmt"

	rpkg "gopkg.in/redis.v4"
)

func ConnectRedis(host string, port int) (client *rpkg.Client, err error) {
	return NewBaseClient(host, port)
}

func NewBaseClient(host string, port int) (client *rpkg.Client, err error) {
	op := &rpkg.Options{
		Addr: fmt.Sprintf("%s:%d", host, port),
	}
	client = rpkg.NewClient(op)

	ping := client.Ping()
	if ping.Err() != nil {
		err = errors.New("faild to connect, reason:" + ping.Err().Error())
		return
	}

	// client.ConnPool.(*rpkg.MultiConnPool).MaxCap = 20
	return
}
