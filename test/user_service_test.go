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

	userSrv := service.NewTUserService(repository.GormDB)
	userInfo, _ := userSrv.GetOne()
	log.Println(userInfo)

}
