package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ckeyer/commons/config"
	"github.com/ckeyer/commons/util"
	"github.com/spf13/viper"
	"time"
)

func init() {
	config.Init("aa")
}

func main() {
	log.Debug("start ")
	log.Info(viper.GetString("url"))
	log.Info(viper.GetString("author.name"))

	var u User
	err := viper.UnmarshalKey("author", &u)
	if err != nil {
		log.Error(err)
		return
	}
	log.Info(u)

	closeAll := make(chan int)
	go func() {
		for {
			select {
			case <-time.Tick(time.Second):
				log.Info("one second gone...")
			case <-closeAll:
				log.Debug("over")
				return
			}
		}
	}()

	util.WaitForExit(true, closeAll)
}

type User struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email`
}
