package model

import (
	"github.com/zooyer/gotpl/common/consts"
	"github.com/zooyer/gotpl/dto"
	"github.com/zooyer/miskit/micro"
)

type Test struct {
	micro.ModelBase
	dto.TestCreateReq
	micro.ModelExtra
}

func (Test) TableName() string {
	return consts.TableNameTest
}
