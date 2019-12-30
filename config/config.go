package config

import (
	"encoding/json"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-cli/init/log"
	"github.com/spf13/viper"
	"time"
)

//Package used
//viper https://github.com/spf13/viper

//GlobalConfig for global configuration
type GlobalConfig struct {
	DB  DBConfig     `json:"db"`
	Sys SystemConfig `json:"sys"`
	JWT JWT          `json:"jwt"`
}

//JWT object to control jwt
type JWT struct {
	HeaderKey  string `json:"headerKey"`
	Secret     string `json:"secret"`
	ExpireTime time.Duration    `json:"expireTime"`
	Issuer     string `json:"issuer"`
}

//SystemConfig manage system settings
type SystemConfig struct {
	IP   string `json:"ip"`
	Port int    `json:"port"`
}

//DBConfig manage DB connect
type DBConfig struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Path         string `json:"path"`
	DBname       string `json:"dbname"`
	Config       string `json:"config"`
	Type         string `json:"type"`
	MaxOpenConns int    `json:"maxOpenConns"`
	MaxIdleConns int    `json:"maxIdleConns"`
}

//Config Global config of DB
var Config GlobalConfig

func init() {
	v := viper.New()
	initConfig(v)
	err := v.ReadInConfig()
	if err != nil {
		panic("check out whether contain resource/application.json file ")
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Log.Infoln("config file has changed", e.Name)
	})
	setDefaultValue(&Config)
	if err := v.Unmarshal(&Config); err != nil {
		log.Log.Errorf("unmarshal application.json fail, error:%s\n", err)
	}
	val, _ := json.MarshalIndent(&Config, " ", "    ")
	log.Log.Infof("start with global config : \n%v\n", string(val))
}

func initConfig(v *viper.Viper) {
	v.SetConfigName("application")
	v.AddConfigPath("resource")
	v.SetConfigType("json")
}

func setDefaultValue(c *GlobalConfig) {
	c.DB.MaxIdleConns = 10
	c.DB.MaxOpenConns = 50
}
