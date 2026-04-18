package conf

import (
	"strings"

	"github.com/spf13/viper"
	"github.com/zooyer/miskit/log"
)

// Viper 全局配置
var Viper = viper.New()

var (
	// env 运行环境
	env string

	// Log 日志配置
	Log log.Config

	// Addr 服务监听地址
	Addr string

	// Jwt Token认证
	Jwt struct {
		Secret string
		Expire int
	}

	// DB 数据库配置
	DB struct {
		Dialect string
		Args    string
		Config  any
	}

	// Cookie 会话配置
	Cookie struct {
		Expires int
	}

	// Wechat 微信信息
	Wechat struct {
		AppID     string `mapstructure:"app_id"`
		AppSecret string `mapstructure:"app_secret"`
	}

	// Douyin 抖音信息
	Douyin struct {
		AppID     string `mapstructure:"app_id"`
		AppSecret string `mapstructure:"app_secret"`
	}
)

var addr = map[string]any{
	"env":       &env,
	"log":       &Log,
	"addr":      &Addr,
	"Jwt":       &Jwt,
	"db":        &DB,
	"db.config": &DB.Config,
	"cookie":    &Cookie,
	"wechat":    &Wechat,
}

func unmarshal() (err error) {
	for key, addr := range addr {
		if err = Viper.UnmarshalKey(key, addr); err != nil {
			return
		}
	}

	return
}

// Init 初始化配置文件
func Init(filename, defaultConfig string) {
	var (
		err    error
		reader = strings.NewReader(defaultConfig)
	)

	Viper.SetConfigType("yaml")

	if err = Viper.ReadConfig(reader); err != nil {
		panic(err)
	}

	if err = unmarshal(); err != nil {
		panic(err)
	}

	Viper.SetConfigFile(filename)

	if err = Viper.ReadInConfig(); err != nil {
		return
	}

	if err = unmarshal(); err != nil {
		panic(err)
	}
}
