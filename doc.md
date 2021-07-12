# dragon 文档

## 简介

一个用 golang 编写的简单易用,且适用于大型项目的web服务端框架

## 目录结构

```tree
.
├── Dockerfile
├── LICENSE
├── README.md
├── bin
│   ├── fswatch
│   └── fswatch.exe
├── constant -------------------------------------------// 常量定义
│   └── constant.go
├── core    // 核心代码
│   └── dragon
│       ├── conf ---------------------------------------// 配置文件，开发，测试，生产相关配置及加载
│       ├── dlogger ------------------------------------// 常用跟踪日志方法
│       ├── dnacos -------------------------------------// 常用跟踪日志方法
│       ├── dragon.go ----------------------------------// http 服务启动器
│       ├── dragoncrypto -------------------------------// 加密方法
│       ├── dragonzipkin -------------------------------// 封装文件压缩方法
│       ├── dredis -------------------------------------// 封装redis库
│       ├── http_context.go ----------------------------// 定义返回数据
│       ├── init.go ------------------------------------// 初始化框架
│       ├── router.go ----------------------------------// 路由
│       ├── tracker ------------------------------------// 链路追踪
│       └── upload.go ----------------------------------// 文件上传
├── docker-compose.yml
├── domain
│   ├── entity  ----------------------------------------// 定义库表结构体
│   │   └── user_entity.go
│   ├── repository -------------------------------------// 仓储层
│   │   ├── base_repository.go
│   │   └── user_repository.go
│   └── service ----------------------------------------// 服务层,用于常用逻辑封装
│       └── user_service.go
├── dto ------------------------------------------------// 定义输入和输入数据结构体
│   └── dto.go
├── go.mod 
├── go.sum
├── handler --------------------------------------------// 事件处理层
│   └── user_handler.go
├── httpclient 
│   └── httpclient.go ----------------------------------// http客户端，封装get,post,put,delete等方法
├── main.go
├── middleware -----------------------------------------// 请求拦截器
│   ├── after_req.go
│   ├── before_req.go
│   └── logger.go
├── outview
│   └── product_view.go --------------------------------// 给客户端定义的输出数据
├── release
│   ├── conf -------------------------------------------// 服务配置文件
│   │   ├── dev.yml
│   │   ├── prod.yml
│   │   └── test.yml
│   └── log --------------------------------------------// 日志文件
│       ├── 2021-04-09.log
│       ├── log
│       └── nacos
├── router ---------------------------------------------// 路由
│   └── router.go 
├── task -----------------------------------------------// 任务
│   └── task.go 
├── test -----------------------------------------------// 单元测试
│   ├── base_repository_test.go
│   ├── dlogger_test.go
│   ├── dredis_test.go
│   ├── httpclient_test.go
│   ├── kafka_test.go
│   ├── nacos_test.go
│   ├── product_service_test.go
│   ├── rabbitmq_test.go
│   └── tools_test.go
├── tools ----------------------------------------------// 一些常用工具
│   ├── dmongo
│   │   └── mongo.go
│   ├── kafka
│   │   └── kafka.go
│   ├── rabbitmq
│   │   └── rabbitmq.go
│   ├── tools.go
│   ├── top.go
│   └── uuid.go
└── vendor ---------------------------------------------// 第三方库
```

## 快速开始

### 配置数据库

### 创建表数据结构

```yml
# project_dir/core/dragon/conf/config/dev.yml
database:
  mysql:
    master:
      host: 127.0.0.1
      port: 3306
      user: root
      password: 123456
      database: travel
      charset: utf8mb4,utf8
      timeout: 3s  #connect timeout time
      maxidle: 20  #maxIdle connections, db will at least serve 20 idle connections; if you set 0, that means no limit。
      maxconn: 100  #maxConn connections, db will have 40 connections limit; if you set 0, that means no limit。
```

### 建表

```sql
CREATE TABLE `t_user`
(
    `user_id`     bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户id',
    `user_nick`   varchar(64) NOT NULL DEFAULT '' COMMENT '用户昵称',
    `user_mobile` varchar(64) NOT NULL DEFAULT '' COMMENT '用户手机号',
    `create_time` datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`user_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户表';
```

```go
// Package entity 创建我们的英雄表数据结构
package entity

type UserEntity struct {
	UserId     int64 `gorm:"primaryKey;AUTO_INCREMENT"`
	UserNick   string
	UserMobile string
	CreateTime string
	UpdateTime string
}

func (UserEntity) TableName() string {
	return "t_user"
}
```

### 创建仓储层

```go
package repository

import (
	"dragon/domain/entity"
	"gorm.io/gorm"
)

// IHeroRepository 定义接口
type IHeroRepository interface {
	IBaseRepository
}

// HeroRepository 定义仓储结构体
type HeroRepository struct {
	BaseRepository
}

// NewUserRepository 定义仓储初始化方法
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		BaseRepository: BaseRepository{
			TableName: entity.UserEntity{}.TableName,
			MysqlDB:   db,
		},
	}
}
```

### 创建服务层

```go
package service

import (
	"dragon/core/dragon/dlogger"
	"dragon/domain/entity"
	"dragon/domain/repository"
	"errors"
	"github.com/go-dragon/erro"
	"gorm.io/gorm"
)

// IUserService 创建服务层接口
type IUserService interface {
}

// UserService 创建服层务结构体
type UserService struct {
	UserRepository repository.IUserRepository
	TxConnDB       *gorm.DB
}

// NewUserService 定义服务初始化方法
func NewUserService(txConnDB *gorm.DB) IUserService {
	return &UserService{
		UserRepository: repository.NewUserRepository(txConnDB),
		TxConnDB:       txConnDB,
	}
}

// GetOneUser AddOne 添加一个英雄
func (u *UserService) GetOneUser(conds []map[string]interface{}, cols string) (*entity.UserEntity, error) {
	var userInfo entity.UserEntity
	res := u.UserRepository.GetOne(&userInfo, conds, cols, "")
	return &userInfo, res.Error
}
```

### 创建handler

```go
package handler

import (
	"dragon/core/dragon"
	"dragon/core/dragon/dlogger"
	"dragon/core/dragon/dredis"
	"dragon/domain/entity"
	"dragon/domain/repository"
	"dragon/domain/service"
	"dragon/handler/reqdata"
	"dragon/tools"
	"encoding/json"
	"github.com/go-dragon/util"
	"github.com/go-dragon/validator"
	validateStruct "github.com/go-playground/validator"
	"net/http"
	"strconv"
	"time"
)

// IUserHandler 定义 handler 接口
type IUserHandler interface {
}

// UserHandler 定义 handler 结构体
type UserHandler struct {
}

func (u *UserHandler) GetOneUser(ctx *dragon.HttpContext) {
	// 初始化req
	var reqData reqdata.UserReq
	err := ctx.BindRequestJsonToStruct(&reqData)
	if err != nil {
		ctx.Json(&dragon.Output{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
			Data: err,
		}, http.StatusBadRequest)
		return
	}
	// 请求参数数据校验
	v := validateStruct.New()
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
		{"member_id": reqData.FirstName},
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
```

### 定义 router

```go
package router

func init() {
	UserHandler = &handler.IUserHandler{}
	// ---------------------------- User相关 -----------------------------------
	dRouter.POST("/getOneUser", UserHandler.GetOneUser)
	// ---------------------------- User相关 -----------------------------------
}
```

### 编译
将编译文件放入 projectDir/release 下运行查看返回结果