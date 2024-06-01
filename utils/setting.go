package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool

	RedisHost string
	RedisPort string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	LoadServer(file)
	LoadData(file)
	LoadMinio(file)
	LoadRedis(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":8088")
	JwtKey = file.Section("server").Key("JwtKey").MustString("89js82js72")
}

func LoadData(file *ini.File) {
	DbHost = file.Section("database").Key("DbHost").MustString("47.121.133.182")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("demo")
	DbName = file.Section("database").Key("DbName").MustString("wedding")
}

func LoadMinio(file *ini.File) {
	Endpoint = file.Section("minio").Key("Endpoint").MustString("47.121.133.182:9000")
	AccessKeyID = file.Section("minio").Key("AccessKeyID").MustString("admin")
	SecretAccessKey = file.Section("minio").Key("SecretAccessKey").MustString("12345678")
	UseSSL = file.Section("minio").Key("UseSSL").MustBool(false)
}

func LoadRedis(file *ini.File) {
	RedisHost = file.Section("redis").Key("RedisHost").MustString("47.121.133.182")
	RedisPort = file.Section("redis").Key("RedisPort").MustString("6379")
}
