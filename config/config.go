package config

import (
	"fmt"
	"os"
	"strconv"
)

var (
	Port  string
	Logs  LogsConfig
	Mysql MysqlConfig
	CertFolderPath string
)

type LogsConfig struct {
	Reporter   bool
	Json       bool
	MaxSize    int
	MaxBackUps int
	MaxAge     int
}

type MysqlConfig struct {
	Database string
	Host     string
	Port     int
	User     string
	Password string
}

func InitConfig() {
	Port = os.Getenv("PORT")
	Logs = initLogsConfig()
	Mysql = initMysqlConfig()
	CertFolderPath = os.Getenv("CERT_FOLDER_PATH")

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

func initMysqlConfig() MysqlConfig {
	return MysqlConfig{
		Database: os.Getenv("MYSQL_DATABASE"),
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     strEnvToInt("MYSQL_PORT"),
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
	}
}

func SetEnvLocal() {

	os.Setenv("MYSQL_DATABASE", "doko")
	os.Setenv("MYSQL_HOST", "localhost")
	os.Setenv("MYSQL_PORT", "3306")
	os.Setenv("MYSQL_USER", "user")
	os.Setenv("MYSQL_PASSWORD", "password")

	os.Setenv("PORT", "8000")

	os.Setenv("LOG_REPORTER", "true")
	os.Setenv("LOG_JSON", "false")
	os.Setenv("LOG_MAX_SIZE", "100")
	os.Setenv("LOG_MAX_BACKUPS", "1")
	os.Setenv("LOG_MAX_AGE", "7")
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
