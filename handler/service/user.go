package service

import (
	"context"

	"github.com/Lilong-maker/zhokao/basic/config"
	__ "github.com/Lilong-maker/zhokao/basic/proto"
	"github.com/Lilong-maker/zhokao/handler/model"
)

<<<<<<< HEAD
=======

import (
	"context"
	"zhokao/basic/config"
	"zhokao/handler/model"


)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	__.StreamGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) GoodAdd(_ context.Context, in *__.GoodAddReq) (*__.GoodAddResp, error) {

	var good model.Show
	err := good.FindGood(config.DB, in.Name)
	if err != nil {
		return &__.GoodAddResp{
			Code: 400,
			Msg:  "商品不存在",
		}, nil
	}
	good.Name = in.Name
	good.Price = float32(in.Price)
	good.Num = int(in.Num)
	err = good.GoodsAdd(config.DB)
	if err != nil {
		return &__.GoodAddResp{
			Code: 400,
			Msg:  "商品添加失败",
		}, nil
	}
	return &__.GoodAddResp{
		Code: 200,
		Msg:  "商品添加成功",
	}, nil
}

// SayHello implements helloworld.GreeterServer
func (s *Server) GoodUpdate(_ context.Context, in *__.GoodUpdateReq) (*__.GoodUpdateResp, error) {
	var good model.Show
	err := good.FIndGoods(config.DB, in.Id)
	if err != nil {
		return &__.GoodUpdateResp{
			Code: 400,
			Msg:  "商品不存在",
		}, nil
	}
	goods := model.Show{
		Name:  good.Name,
		Price: good.Price,
		Num:   good.Num,
	}
	err = goods.GoodsUpdate(config.DB, in.Id)
	if err != nil {
		return &__.GoodUpdateResp{
			Code: 200,
			Msg:  "修改成功",
		}, nil
	}
	return &__.GoodUpdateResp{
		Code: 200,
		Msg:  "修改成功",
	}, nil
}

// SayHello implements helloworld.GreeterServer
func (s *Server) GoodList(_ context.Context, in *__.GoodListReq) (*__.GoodListResp, error) {
	var good model.Show
	err := good.FIndGoods(config.DB, in.Id)
	if err != nil {
		return &__.GoodListResp{
			Code: 400,
			Msg:  "列表查询失败",
		}, nil
	}
	return &__.GoodListResp{
		Code:  200,
		Msg:   "列表查询成功",
		Name:  good.Name,
		Price: uint32(good.Price),
		Num:   int32(good.Num),
	}, nil
}

// SayHello implements helloworld.GreeterServer
func (s *Server) GoodDelete(_ context.Context, in *__.GoodDeleteReq) (*__.GoodDeleteResp, error) {
	var good model.Show
	err := good.FIndGoods(config.DB, in.Id)
	if err != nil {
		return &__.GoodDeleteResp{
			Code: 400,
			Msg:  "商品不存在",
		}, nil
	}
	err = good.DeleteGood(config.DB, in.Id)
	if err != nil {
		return &__.GoodDeleteResp{
			Code: 400,
			Msg:  "删除失败",
		}, nil
	}
	return &__.GoodDeleteResp{
		Code: 200,
		Msg:  "删除成功",
	}, nil
}
>>>>>>> local_0122
