package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ckeyer/commons/config"
	"github.com/spf13/viper"
)

func init() {
	config.Init("aa")
}

func main() {
	log.Debug("start ")
	log.Info(viper.GetString("url"))
}
