package lib

import (
	"fmt"

	redis "gopkg.in/redis.v3"
)

type RedisConfig struct {
	Host     string
	Port     string
	Password string
}

func GetRedis(r *RedisConfig) (*redis.Client, error) {
	connstr := fmt.Sprintf("%s:%s", r.Host, r.Port)
	redis_cli = redis.NewClient(&redis.Options{
		Addr:     connstr,
		Password: r.Password,
		DB:       0,
	})
	_, err := redis_cli.Ping().Result()
	if err != nil {
		return nil, err
	}
	return redis_cli, nil
}
