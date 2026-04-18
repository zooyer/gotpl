package service

import (
	"context"

	"github.com/zooyer/gotpl/common/errno"
	"github.com/zooyer/gotpl/dao"
	"github.com/zooyer/gotpl/dto"
	"github.com/zooyer/gotpl/model"
	"github.com/zooyer/miskit/micro"
	"github.com/zooyer/miskit/utils"
)

type _test_ struct {
	micro.Service[*model.User, model.Test, *model.Test]
}

var Test _test_

func (t _test_) Get(ctx context.Context, id int64) (*model.Test, error) {
	return t.Service.Get(ctx, micro.Getter{
		Dao:   dao.Test,
		Errno: errno.TestGetError,
		Equal: micro.Equal{"id": id},
	})
}

func (t _test_) List(ctx context.Context, query micro.Query) (*micro.Result[model.Test, *model.Test], error) {
	return t.Service.List(ctx, micro.Lister{
		Dao:   dao.Test,
		Errno: errno.TestListError,
		Query: query,
	})
}

func (t _test_) Create(ctx context.Context, req dto.TestCreateReq) (*model.Test, error) {
	return t.Service.Create(ctx, micro.Creator[model.Test, *model.Test]{
		Dao:   dao.Test,
		Errno: errno.TestCreateError,
		Equal: micro.Equal{"name": req.Name},
		Model: model.Test{TestCreateReq: req},
	})
}

func (t _test_) Update(ctx context.Context, req dto.TestUpdateReq) (*model.Test, error) {
	return t.Service.Update(ctx, micro.Updater{
		Dao:    dao.Test,
		Errno:  errno.TestUpdateError,
		Equal:  micro.Equal{"id": req.ID.ID},
		Update: utils.StructToMap(req.TestCreateReq, "json"),
	})
}

func (t _test_) Delete(ctx context.Context, req dto.TestDeleteReq) (*model.Test, error) {
	return t.Service.Delete(ctx, micro.Deleter{
		Dao:   dao.Test,
		Errno: errno.TestDeleteError,
		Equal: micro.Equal{"id": req.ID},
	})
}
