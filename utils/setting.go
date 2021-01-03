package utils

import (
	"github.com/spf13/viper"
	"os"
)

// 读取配置文件config

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	Db         string
	DbHost     string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassWord string

	AccessKey  string
	SecretKey  string
	Bucket     string
	QiniuSever string
)

func init() {
	//获取项目的执行路径
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	config := viper.New()

	config.AddConfigPath(path + "/config") //设置读取的文件路径
	config.SetConfigName("config")         //设置读取的文件名
	config.SetConfigType("yaml")           //设置文件的类型
	//尝试进行配置读取
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}

	// 读取配置数据
	AppMode = config.GetString("server.AppMode")
	HttpPort = config.GetString("server.HttpPort")
	JwtKey = config.GetString("server.JwtKey")

	Db = config.GetString("database.Db")
	DbHost = config.GetString("database.DbHost")
	DbPort = config.GetString("database.DbPort")
	DbName = config.GetString("database.DbName")
	DbUser = config.GetString("database.DbUser")
	DbPassWord = config.GetString("database.DbPassWord")

	AccessKey = config.GetString("database.AccessKey ")
	SecretKey = config.GetString("database.SecretKey ")
	Bucket = config.GetString("database.Bucket")
	QiniuSever = config.GetString("database.QiniuSever")

}
