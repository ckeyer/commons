package config

import (
	"fmt"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

var appName = "ck"

func Init(name ...string) {
	if len(name) == 1 {
		appName = name[0]
	}

	viper.SetEnvPrefix(appName)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	configPath := os.Getenv("CFG_PATH")
	if configPath != "" {
		viper.AddConfigPath(configPath)
	} else {
		viper.AddConfigPath("./")
		viper.AddConfigPath("conf/")
	}

	viper.SetConfigName(appName)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error when reading %s config file: %s\n", appName, err))
	}

	switch viper.GetString("logging.level") {
	case "debug", "d":
		log.SetLevel(log.DebugLevel)
	case "info", "i":
		log.SetLevel(log.InfoLevel)
	case "warning", "w":
		log.SetLevel(log.WarnLevel)
	case "error", "e":
		log.SetLevel(log.ErrorLevel)
	default:
	}

}
