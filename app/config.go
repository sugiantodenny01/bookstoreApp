package app

import "github.com/tkanos/gonfig"

type Config struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
	Host        string
}

func GetConfig() Config {
	conf := Config{}
	gonfig.GetConf("app/config.json", &conf)
	return conf
}
