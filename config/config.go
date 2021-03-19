package config

import (
	"fmt"
	"os"
	"strconv"
)

var (
	Port string
	Logs LogsConfig
)

type LogsConfig struct {
	Reporter   bool
	Json       bool
	MaxSize    int
	MaxBackUps int
	MaxAge     int
}

func InitConfig() {
	Port = os.Getenv("PORT")
	Logs = initLogsConfig()
}

func initLogsConfig() LogsConfig {
	return LogsConfig{
		Reporter:   strEnvToBool("LOG_REPORTER"),
		Json:       strEnvToBool("LOG_JSON"),
		MaxSize:    strEnvToInt("LOG_MAX_SIZE"),
		MaxBackUps: strEnvToInt("LOG_MAX_BACKUPS"),
		MaxAge:     strEnvToInt("LOG_MAX_AGE"),
	}
}

func strEnvToInt(envString string) int {
	sValue := os.Getenv(envString)
	i, err := strconv.Atoi(sValue)
	if err != nil {
		panic(fmt.Errorf("can't parse env '%s' with value '%s' %s", envString, sValue, err))
	}
	return i
}

func strEnvToBool(envString string) bool {
	sValue := os.Getenv(envString)
	b, err := strconv.ParseBool(sValue)
	if err != nil {
		panic(fmt.Errorf("can't parse '%s' %s", envString, err))
	}
	return b
}
