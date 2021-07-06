package service

import (
	"dragon/domain/entity"
	"dragon/domain/repository"
	"errors"
	"github.com/go-dragon/util"
	"gorm.io/gorm"
	"time"
)

//  UserService interface
type IUserService interface {
	GetOne() (*entity.UserEntity, error)
	TransactionTest() error
}

type UserService struct {
	UserRepository repository.IUserRepository
	TxConnDB       *gorm.DB //当前服务所用的一条DB连接，用于事务处理
}

func NewUserService(txConnDB *gorm.DB) IUserService {
	return &UserService{
		UserRepository: repository.NewUserRepository(txConnDB),
		TxConnDB:       txConnDB,
	}
}

// 获取一条 todo 测试事务
func (u *UserService) GetOne() (*entity.UserEntity, error) {
	var user entity.UserEntity
	var conditions = []map[string]interface{}{
		{"user_id = ?": 1},
		//{"create_time <= ?": "2019-08-01"},
	}
	res := u.UserRepository.GetOne(&user, conditions, "*", "")
	if repository.HasSeriousError(res) {
		return &user, res.Error
	}
	return &user, nil
}

// test transaction easy demo
func (u *UserService) TransactionTest() error {
	userInfo := entity.UserEntity{
		UserNick:   util.RandomStr(6),
		UserMobile: "",
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
		UpdateTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	err1 := u.UserRepository.Add(&userInfo)
	userInfo = entity.UserEntity{
		UserNick:   util.RandomStr(4),
		UserMobile: "",
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
		UpdateTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	err2 := u.UserRepository.Add(&userInfo)
	err1 = errors.New("manual error")
	if err1 != nil || err2 != nil {
		u.TxConnDB.Rollback()
		return errors.New("data write fail")
	}
	u.TxConnDB.Commit()
	return nil
}
