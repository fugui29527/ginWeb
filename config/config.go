package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"path/filepath"
)

var AppConfig *Config

//  系统信息
type appInfo struct {
	Port int `yaml:"port"`
}

/**
  db信息
*/
type dbInfo struct {
	Adress   string `yaml:"adress"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	PassWord string `yaml:"password"`
	DbName   string `yaml:"dbName"`
	Showsql  bool   `yaml:"showsql"`
}
type redisInfo struct {
	Addr     string `yaml:"addr"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type logInfo struct {
	Level    string `yaml:"level"`
	Encoding string `yaml:"encoding"`
	Path     string `yaml:"path"`
}

//配置信息
type Config struct {
	AppInfo   appInfo   `yaml:"appInfo"`
	DbInfo    dbInfo    `yaml:"dbInfo"`
	RedisInfo redisInfo `yaml:"redis"`
	LogInfo   logInfo   `yaml:"log"`
}

func init() {
	log.Printf("=======初始化配置文件===============")
	//初始化配置文件
	initConfig()
}

func initConfig() {
	//获取当前项目路劲
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		panic("获取项目路径失败!")
	}
	path := filepath.Join(dir, "config.yaml")
	AppConfig = &Config{}
	if f, err := os.Open(path); err != nil {
		log.Printf("=======初始化配置文件失败:", err)
		panic("初始化配置文件失败")
	} else {
		yaml.NewDecoder(f).Decode(AppConfig)
	}
}
