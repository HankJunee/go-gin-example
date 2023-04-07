package config

import (
	"gopkg.in/ini.v1"
	"log"
)

const (
	configFilePath = "./config/config.ini"
)

var (
	DBConfig  = Database{}
	AppConfig = App{}
)

type Database struct {
	DbHost      string
	DbPort      string
	UserName    string
	Password    string
	DBName      string
	TimeZone    string
	TablePrefix string
}

type App struct {
	RunPort string
}

func init() {
	cfg, err := ini.Load(configFilePath)

	if err != nil {
		log.Fatalln("配置文件读取失败，请检查配置文件路径是否正确")
	}

	err = cfg.Section("Database").MapTo(&DBConfig)

	if err != nil {
		log.Fatalln("配置文件解析失败，请检查配置文件路径是否正确")
	}

	err = cfg.Section("App").MapTo(&AppConfig)

	if err != nil {
		log.Fatalln("配置文件解析失败，请检查配置文件路径是否正确")
	}
}
