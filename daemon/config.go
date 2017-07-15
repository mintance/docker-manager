package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"github.com/Sirupsen/logrus"
)

type Config struct {
	System struct {
		Host string `yaml:"listen_host"`
		Port string `yaml:"listen_port"`
	} `yaml:"system"`
}

const config_path  = `config/config.yaml`

func getConfig() *Config {

	config := Config{}

	logrus.Info("Reading config file: " + config_path)

	data, err := ioutil.ReadFile(config_path)

	if err != nil {
		logrus.Fatal("Config open error: ", err)
	}

	err = yaml.Unmarshal(data, &config)

	if err != nil {
		logrus.Fatal("Config read & unmarshal error: ", err)
	}

	return &config
}