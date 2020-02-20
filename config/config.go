package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	AppName string
	RunMode string
	Listen string
	StaticPath string
	MysqlDsn string

	VoiceNoteUrl string
}

var (
	cnf Config
)

func init() {
	viper.SetConfigName("app")
	viper.AddConfigPath("./conf")
	err := viper.ReadInConfig()

	if err != nil {
		panic("Error load config!")
	}

	viper.Unmarshal(&cnf)
}

func AppConfig() Config {
	return cnf
}

func IsProd() bool {
	return cnf.RunMode == "prod"
}