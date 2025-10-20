package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	System SystemConfig `yaml:"system"`
	Mysql  MysqlConfig  `yaml:"mysql"`
}

type SystemConfig struct {
	Domain      string `yaml:"domain"`
	Version     string `yaml:"version"`
	Env         string `yaml:"env"`
	HttpPort    string `yaml:"HttpPort"`
	Host        string `yaml:"Host"`
	UploadModel string `yaml:"UploadModel"`
}

type MysqlConfig struct {
	Default DbConfig `yaml:"default"`
}

type DbConfig struct {
	Dialect  string `yaml:"dialect"`
	DbHost   string `yaml:"dbHost"`
	DbPort   string `yaml:"dbPort"`
	DbName   string `yaml:"dbName"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
}

var AppCfg = &Config{}

func InitConfig() error {
	// 读取配置文件
	data, err := os.ReadFile("config/config.yaml")
	if err != nil {
		return fmt.Errorf("failed to read config file: %v", err)
	}

	// 解析YAML
	err = yaml.Unmarshal(data, AppCfg)
	if err != nil {
		return fmt.Errorf("failed to parse config file: %v", err)
	}

	return nil
}
