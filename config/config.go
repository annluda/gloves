package config

import (
	"github.com/spf13/viper"
)

const (
	// RunmodeDebug -
	RunmodeDebug = "debug"
	// RunmodeRelease -
	RunmodeRelease = "release"
	// RunmodeTest -
	RunmodeTest = "test"

	// 配置文件路径
	configFilePath = "./config.yaml"
	// 日志文件路径
	logFilePath = "logs/gloves.log"
	// 配置文件格式
	configFileType = "yaml"
)

var (
	// AppConfig 应用配置
	AppConfig *appConfig
	// DBConfig 数据库配置
	DBConfig *dbConfig
	// MailConfig 邮件配置
	MailConfig *mailConfig
)

// InitConfig 初始化配置
func InitConfig() {
	// 初始化 viper 配置
	viper.SetConfigFile(configFilePath)
	viper.SetConfigType(configFileType)

	//if err := viper.ReadInConfig(); err != nil {
	//  panic(fmt.Sprintf("读取配置文件失败，请检查: %v", err))
	//}

	// 初始化 app 配置
	AppConfig = newAppConfig()
	// 初始化数据库配置
	//DBConfig = newDBConfig()
	// 初始化邮件配置
	//MailConfig = newMailConfig()

}
