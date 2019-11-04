package config

import (
	yaml2 "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// DB 数据库配置
type Db struct {
	EnableLog          bool   `yaml:"enable_log"`
	Dialect            string `yaml:"dialect"`
	Host               string `yaml:"host"`
	User               string `yaml:"user"`
	PassWd             string `yaml:"pass" json:"pass"`
	Db                 string `yaml:"db"`
	MaxOpenConnections int    `yaml:"max_open_connections"`
	MaxIdleConnections int    `yaml:"max_idle_connections"`
}

// Service 服务端配置
type Service struct {
	Mode        string `yaml:"mode"`
	Port        string `yaml:"port"`
	TCPPort     string `yaml:"tcp_port"`
	ServiceUrl  string `yaml:"service_url"`
	RpcUrl      string `yaml:"rpc_url"`
	MaxPageSize int    `yaml:"max_page_size"`
}

// redis
type Redis struct {
	Host   string `yaml:"host" json:"host"`
	PassWd string `yaml:"pass" json:"pass_wd"`
	Db     int    `yaml:"db" json:"db"`
}

type LogConfig struct {
	Path string `yaml:"path"`
}

type WssConfig struct {
	HeartbeatTime int64 `yaml:"heart_beat_time"`
}

// Config 配置
type Config struct {
	Service Service   `yaml:"service"`
	DB      Db        `yaml:"db"`
	Redis   Redis     `yaml:"redis"`
	Log     LogConfig `yaml:"log"`
	Wss     WssConfig `yaml:"wss"`
}

func GetDb() Db {
	return config.DB
}

func GetService() Service {
	return config.Service
}

func GetRedis() Redis {
	return config.Redis
}

func GetWss() WssConfig {
	return config.Wss
}

func GetLog() LogConfig {
	return config.Log
}

var config Config

func InitConfig(path string) {
	pathStr, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	yamlFile, err := ioutil.ReadFile(pathStr + path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml2.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Printf("Unmarshal: %v", err)
	}

}
