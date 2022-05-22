package conf

import (
	"github.com/spf13/viper"
	"log"
)

// 声明一个全局的配置对象
var Config = new(ConfigStruct)

// 声明一个全局的配置对象
type ConfigStruct struct {
	// MySQL数据库配置
	*DatabaseConfig `mapstructure:"database"`
	*LoggerConfig   `mapstructure:"logger"`
	*CacheConfig    `mapstructure:"cache"`
	*JwtConfig      `mapstructure:"jwt"`
}

// 初始化配置
func InitConfig() {
	log.Println("初始化配置")
	// viper 配置文件
	viper.SetConfigFile("./config/config.yaml")
	// 读取配置
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("Fatal error config file: %s \n", err)
		return
	}
	// 解析配置文件
	if err := viper.Unmarshal(Config); err != nil {
		log.Fatalln("viper.Unmarshal failed, err:%v\n", err)
		return
	}
	log.Println("配置初始化完成")
}

// 获取配置
func GetConfig() *ConfigStruct {
	return Config
}
