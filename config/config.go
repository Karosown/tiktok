package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

/**
定义结构体
*/

type Config struct {
	Svc server   `yaml:"server"`   //server配置
	DB  database `yaml:"database"` //数据库配置
}

type server struct {
	Port     int    `yaml:"port"`
	RunMode  string `yaml:"runMode"`
	LogLevel string `yaml:"logLevel"`
}

type database struct {
	Type            string `yaml:"type"`
	Host            string `yaml:"host"`
	Port            int    `yaml:"port"`
	UserName        string `yaml:"username"`
	Password        string `yaml:"password"`
	DbName          string `yaml:"dbname"`
	MaxIdleConn     int    `yaml:"max_idle_conn"`
	MaxOpenConn     int    `yaml:"max_open_conn"`
	ConnMaxLifetime string `yaml:"conn_max_lifetime"`
}

/**
相当于重写toString方法
*/

func (c Config) String() string {
	return fmt.Sprintf("%v\n%v\n", c.Svc, c.DB)
}

func (s server) String() string {
	return fmt.Sprintf("server : \n"+
		"\tport : %v \n"+
		"\tRunMode : %v", s.Port, s.RunMode)
}

func (m database) String() string {
	return fmt.Sprintf("database : \n"+
		"\ttype : %v \n"+
		"\thost : %v \n"+
		"\tport : %v \n"+
		"\tusername : %v \n"+
		"\tpassword : %v \n"+
		"\tdbname : %v \n"+
		"\tmax_idle_conn : %v \n"+
		"\tmax_open_conn : %v \n"+
		"\tconn_max_lifetime : %v",
		m.Type, m.Host, m.Port, m.UserName, m.Password, m.DbName, m.MaxOpenConn, m.MaxIdleConn, m.ConnMaxLifetime)
}

/**
连接
*/

func InitConfig() *Config {
	dataFile := "config.yml"
	yamlFile, err := os.ReadFile(dataFile)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	c := new(Config)
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Printf("Unmarshal: %v", err)
	}
	log.Printf("load conf success\n %v", c)
	// 绑定到外部可以访问的变量中
	fmt.Printf("config.server: %#v\n", c.Svc)
	fmt.Printf("config.data: %#v\n", c.DB)
	return c
}
