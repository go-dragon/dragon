package handler

import (
	"dragon/core/dragon"
	"dragon/core/dragon/dlogger"
	"dragon/core/dragon/goroutine"
	"dragon/domain/repository"
	"dragon/domain/service"
	"dragon/handler/reqdata"
	"github.com/go-dragon/erro"
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
	gid := goroutine.CurGoroutineID()
	err := ctx.BindPostJson(&userReq)
	if err != nil {
		errs := erro.NewError(err)
		dlogger.Error(errs)
		ctx.Json(&dragon.Output{
			Code: http.StatusBadRequest,
			Msg:  http.StatusText(http.StatusBadRequest),
			Data: gid,
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
