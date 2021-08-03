package service

import (
	"dragon/domain/entity"
	"dragon/domain/repository"
	"errors"
	"gorm.io/gorm"
)

// IUserService interface
type IUserService interface {
	GetOneByUserId(userId int64) (*entity.UserEntity, error)
	TransactionTest() error
	AddOne(userEntity *entity.UserEntity) error
}

type UserService struct {
	UserRepository repository.IUserRepository
}

func NewUserService(txConnDB *gorm.DB) IUserService {
	return &UserService{
		UserRepository: repository.NewUserRepository(txConnDB),
	}
}

func (u *UserService) GetOneByUserId(userId int64) (*entity.UserEntity, error) {
	var user *entity.UserEntity
	user, err := u.UserRepository.GetOneByUserId(userId)
	return user, err
}

// TransactionTest transaction easy demo
func (u *UserService) TransactionTest() error {
	// begin transaction, 如果是多service操作，需要将u.TxConnDB传入其他service
	transDB := repository.GormDB.Begin()
	userSrv := NewUserService(transDB)
	userInfo := entity.UserEntity{
		UserNick:   "superman",
		UserMobile: "11111555556666",
	}
	err := userSrv.AddOne(&userInfo)
	if err != nil {
		transDB.Rollback()
		return errors.New("data write fail")
	}
	transDB.Commit()
	return nil
}

// AddOne insert one data
func (u *UserService) AddOne(userEntity *entity.UserEntity) error {
	err := u.UserRepository.AddOne(userEntity)
	return err
}
