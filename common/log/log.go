package log

import (
	"github.com/zooyer/gotpl/common/conf"
	"github.com/zooyer/miskit/log"
)

var Logger *log.Logger

func Init() {
	var color bool
	if conf.IsDev() || conf.IsTest() {
		color = true
	}

	logger, err := log.New(conf.Log, log.TextFormatter(color))
	if err != nil {
		panic(err)
	}

	Logger = logger
}
