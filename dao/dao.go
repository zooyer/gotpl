package dao

import (
	"context"

	"github.com/zooyer/gotpl/common/conf"
	"github.com/zooyer/gotpl/model"
	"github.com/zooyer/miskit/micro"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB(ctx context.Context) *gorm.DB {
	if conf.IsDev() || conf.IsTest() {
		return db.WithContext(ctx).Debug()
	}

	return db.WithContext(ctx)
}

func Init() {
	var err error
	if db, err = micro.OpenDB(conf.DB.Dialect, conf.DB.Args, nil); err != nil {
		panic(err)
	}

	db = db.Unscoped()

	var ctx = context.Background()
	if err = DB(ctx).AutoMigrate(
		new(model.User),
		new(model.Test),
	); err != nil {
		panic(err)
	}
}

var (
	User = micro.NewDao(DB, new(model.User))
	Test = micro.NewDao(DB, new(model.Test))
)
