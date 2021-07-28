package handler

import (
	"dragon/core/dragon"
	"dragon/domain/repository"
	"dragon/domain/service"
	"dragon/handler/reqdata"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type IUserHandler interface {
	Test(ctx *dragon.HttpContext)
}

// UserHandler 用户处理器
type UserHandler struct {
}

func (u *UserHandler) Test(ctx *dragon.HttpContext) {
	var userReq reqdata.UserReq
	// bind json to struct
	//gid := goroutine.CurGoroutineID()
	err := ctx.BindPostJson(&userReq)
	if err != nil {
		ctx.Json(&dragon.Output{
			Code: http.StatusBadRequest,
			Msg:  http.StatusText(http.StatusBadRequest),
			Data: nil,
		}, http.StatusBadRequest)
		return
	}

	v := validator.New()
	err = v.Struct(&userReq)
	if err != nil {
		ctx.Json(&dragon.Output{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
			Data: err,
		}, http.StatusBadRequest)
		return
	}
	// if something error
	// dlogger.Error("errors")

	// mongodb example
	//mongoRes, err := dmongo.DefaultDB().Collection("c_device_log").InsertOne(context.Background(), bson.M{
	//	"device_name": "golang",
	//})
	//if err != nil {
	//	fmt.Println("mongoErr", err)
	//}
	//objectId := mongoRes.InsertedID.(primitive.ObjectID)
	//fmt.Println("mongoRes", hex.EncodeToString(objectId[:]))

	// mysql example
	//log.Println("reqParams", fmt.Sprintf("%+v", ctx.GetRequestParams()))
	userSrv := service.NewUserService(repository.GormDB) // 如果是事务处理，这个db可以为gorm的begin的db，只能从头传进去🤷‍
	res, _ := userSrv.GetOne()
	//log.Println("err:", err)

	//res := dto.TStructToData(product, []string{"product_id", "product_name", "create_time"})

	output := dragon.Output{
		Code: http.StatusOK,
		Msg:  "ok",
		Data: res,
	}
	ctx.Json(&output, http.StatusOK)
	return
}

// GetOneUser 获取单个用户信息
func (u *UserHandler) GetOneUser(ctx *dragon.HttpContext) {
	// 初始化req
	var reqData reqdata.UserReq
	err := ctx.BindPostJson(&reqData)
	if err != nil {
		ctx.Json(&dragon.Output{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
			Data: err,
		}, http.StatusBadRequest)
		return
	}
	// 请求参数数据校验
	v := validator.New()
	err = v.Struct(&reqData)
	if err != nil {
		ctx.Json(&dragon.Output{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
			Data: err,
		}, http.StatusBadRequest)
		return
	}
	mbSrv := service.NewUserService(repository.GormDB)
	conds := []map[string]interface{}{
		{"member_id = ?": reqData.FirstName},
	}
	mbInfo, err := mbSrv.GetOneUser(conds, "*")
	if err != nil {
		ctx.Json(&dragon.Output{
			Code: http.StatusBadRequest,
			Msg:  "未查询到用户信息",
			Data: err,
		}, http.StatusBadRequest)
		return
	}
	ctx.Json(&dragon.Output{
		Code: http.StatusOK,
		Msg:  "OK",
		Data: mbInfo,
	}, http.StatusOK)
	return
}
