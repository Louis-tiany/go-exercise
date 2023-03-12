/*
* File     : config.go
* Author   : *
* Mail     : *
* Creation : Fri 10 Mar 2023 10:51:44 PM CST
 */
package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

var (
	Config *CommonConfig
)

type CommonConfig struct {
	DbConfig struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DbName   string `yaml:"db_name"`
	}
	External struct {
		HodayUrl string `yaml:"holiday_url"`
		SinaUrl  string `yaml:"sina_url"`
	}
}

func readConf(filename string) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("config file: " + filename + "open error")
	}
	Config = new(CommonConfig)
	err = yaml.Unmarshal(buf, Config)
	if err != nil {
		panic("config file: " + filename + "init error")
	}
}

func init() {
	readConf("./stock.yaml")
}
