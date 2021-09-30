package test

import (
	"dragon/core/dragon"
	"dragon/domain/repository"
	"dragon/domain/service"
	"log"
	"testing"
)

func TestUserService_TransactionTest(t *testing.T) {
	dragon.AppInit()

	userSrv := service.NewUserService(repository.GormDB)
	userInfo, _ := userSrv.GetOne()
	log.Println(userInfo)

}
