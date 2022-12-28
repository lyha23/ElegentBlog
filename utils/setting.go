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

	TencentAccessKey string
	TencentSecretKey string
	TencentBucket    string
	TencentSever     string

	AdminName string
	AdminPwd  string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	LoadServer(file)
	LoadData(file)
	LoadTencentyun(file)
	LoadAdminConfig(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("89js82js72")
}

func LoadData(file *ini.File) {
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("ginblog")
	DbPassWord = file.Section("database").Key("DbPassWord").String()
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}

func LoadTencentyun(file *ini.File) {
	TencentAccessKey = file.Section("tencentyun").Key("AccessKey").String()
	TencentSecretKey = file.Section("tencentyun").Key("SecretKey").String()
	TencentBucket = file.Section("tencentyun").Key("Bucket").String()
	TencentSever = file.Section("tencentyun").Key("Sever").String()
}

func LoadAdminConfig(file *ini.File) {
	AdminName = file.Section("admin").Key("Username").MustString("admin")
	AdminPwd = file.Section("admin").Key("Password").MustString("admin")
}