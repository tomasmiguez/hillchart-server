package configs

import (
    "log"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type DbConfig struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"db_name"`
	Port     string `yaml:"port"`
}

func Db() DbConfig {
	configFile, err := ioutil.ReadFile("configs/database.yaml")
	if err != nil {
		log.Fatal(err)
	}

	dbConfig := DbConfig{}
	err = yaml.Unmarshal(configFile, &dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	return dbConfig
}
