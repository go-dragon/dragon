package test

import (
	"dragon/core/dragon"
	"dragon/domain/repository"
	"dragon/domain/service"
	"testing"
)

func TestUserService_TransactionTest(t *testing.T) {
	dragon.AppInit()

	userSrv := service.NewUserService(repository.GormDB)
	userSrv.TransactionTest()

}
