package lib

import (
	"fmt"

	redis "gopkg.in/redis.v3"
)

func GetRedis(host, port, password string) (*redis.Client, error) {
	connstr := fmt.Sprintf("%s:%s", host, port)
	redis_cli := redis.NewClient(&redis.Options{
		Addr:     connstr,
		Password: password,
		DB:       0,
	})
	_, err := redis_cli.Ping().Result()
	if err != nil {
		return nil, err
	}
	return redis_cli, nil
}
