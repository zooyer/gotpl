package main

import (
	"embed"
	"flag"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/zooyer/gotpl/common/conf"
	"github.com/zooyer/gotpl/common/consts"
	"github.com/zooyer/gotpl/common/log"
	"github.com/zooyer/gotpl/dao"
	"github.com/zooyer/gotpl/handler"
	"github.com/zooyer/miskit/micro"
	"github.com/zooyer/miskit/utils/antd"
	"github.com/zooyer/miskit/utils/webfs"
)

//go:embed web/dist
var static embed.FS

//go:embed conf/dev.yaml
var defaultConfig string

var config = flag.String("config", "conf/conf.dev.yaml", "config file")

func init() {
	flag.Parse()

	conf.Init(*config, defaultConfig)
	log.Init()
	dao.Init()
}

func main() {
	engine := gin.Default()

	// 开发环境
	if conf.IsDev() || conf.IsTest() {
		engine = gin.Default()
	}

	// 会话
	store := cookie.NewStore([]byte(consts.CookieSecret))
	store.Options(sessions.Options{
		Path:     "/",
		Domain:   "",
		MaxAge:   conf.Cookie.Expires,
		Secure:   false,
		HttpOnly: false,
		SameSite: 0,
	})
	engine.Use(sessions.Sessions(consts.SessionName, store))

	// 中间件
	engine.Use(micro.Recover(log.Logger)) // 异常捕获
	engine.Use(micro.Logger(log.Logger))  // 日志记录
	engine.Use(micro.Trace(consts.Name))  // 链路追踪

	// 前端 antd
	efs := webfs.WebEFS("/web/dist", static)
	afs := antd.AntFS(
		efs,
		"",
		"logo.png",
		"logo.svg",
		"favicon.ico",
		"scripts/loading.js",
		"/:route",
		"/:route/:route",
	)
	afs.Mount("/static", engine)

	v1 := engine.Group(fmt.Sprintf("/%s/api/v1", consts.Name))
	{
		test := v1.Group("/test")
		{
			test.GET("", handler.Test.Get)
			test.GET("/list", handler.Test.Get)
			test.POST("/create", handler.Test.Create)
			test.POST("/update", handler.Test.Update)
			test.POST("/delete", handler.Test.Delete)
		}
	}

	if err := engine.Run(conf.Addr); err != nil {
		panic(err)
	}
}
