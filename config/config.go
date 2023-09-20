package config

import (
	"gopkg.in/ini.v1"
	"log"
)

var (
	AppMode  string
	HttpPort string
	BaseUrl  string
	
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
	
	RedisDb     string
	RedisAddr   string
	RedisPw     string
	RedisDbName string
)

// Init 初始化配置
func Init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		log.Panicln("配置文件读取错误，请检查文件路径:", err)
	}
	
	LoadServerConfig(file)
	LoadMySqlConfig(file)
	LoadRedisConfig(file)
	
	log.Println("配置读取成功")
}

// LoadServerConfig 获取service配置
func LoadServerConfig(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
	BaseUrl = file.Section("service").Key("BaseUrl").String()
	
}

// LoadMySqlConfig 读取mySql配置
func LoadMySqlConfig(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassword = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}

// LoadRedisConfig 读取redis配置
func LoadRedisConfig(file *ini.File) {
	RedisDb = file.Section("redis").Key("RedisDb").String()
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
}
