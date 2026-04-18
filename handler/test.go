package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zooyer/gotpl/dto"
	"github.com/zooyer/gotpl/model"
	"github.com/zooyer/gotpl/service"
	"github.com/zooyer/miskit/micro"
)

type _test_ struct {
	micro.Controller
}

var Test _test_

func (t _test_) Get(ctx *gin.Context) {
	var (
		err  error
		req  dto.TestGetReq
		resp *model.Test
	)

	defer func() { t.Response(ctx, resp, err) }()

	if err = t.Bind(ctx, &req); err != nil {
		return
	}

	if resp, err = service.Test.Get(ctx, req.ID); err != nil {
		return
	}
}

func (t _test_) List(ctx *gin.Context) {
	var (
		err    error
		query  micro.Query
		result *micro.Result[model.Test, *model.Test]
	)

	defer func() { t.Response(ctx, result, err) }()

	if err = t.Bind(ctx, &query); err != nil {
		return
	}

	if result, err = service.Test.List(ctx, query); err != nil {
		return
	}
}

func (t _test_) Create(ctx *gin.Context) {
	var (
		err  error
		req  dto.TestCreateReq
		resp *model.Test
	)

	defer func() { t.Response(ctx, resp, err) }()

	if err = ctx.Bind(&req); err != nil {
		return
	}

	if resp, err = service.Test.Create(ctx, req); err != nil {
		return
	}
}

func (t _test_) Update(ctx *gin.Context) {
	var (
		err  error
		req  dto.TestUpdateReq
		resp *model.Test
	)

	defer func() { t.Response(ctx, resp, err) }()

	if err = ctx.Bind(&req); err != nil {
		return
	}

	if resp, err = service.Test.Update(ctx, req); err != nil {
		return
	}
}

func (t _test_) Delete(ctx *gin.Context) {
	var (
		err  error
		req  dto.TestDeleteReq
		resp *model.Test
	)

	defer func() { t.Response(ctx, resp, err) }()

	if err = ctx.Bind(&req); err != nil {
		return
	}

	if resp, err = service.Test.Delete(ctx, req); err != nil {
		return
	}
}
