package test

import (
	"dragon/core/dragon"
	"dragon/domain/repository"
	"dragon/domain/service"
	"testing"
)

func TestProductService_TransactionTest(t *testing.T) {
	dragon.AppInit()

	tx := repository.GormDB.Begin()
	userSrv := service.NewUserService(tx)
	userSrv.TransactionTest()

}
