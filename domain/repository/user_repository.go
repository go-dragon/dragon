package repository

import (
	"dragon/domain/entity"
	"gorm.io/gorm"
)

// all the sql
const (
	GetOneByUserId = `select * from t_user where user_id = ?`
	AddOne         = `insert into t_user (user_nick, user_mobile) values (?, ?)`
)

// IUserRepository user Repo
type IUserRepository interface {
	GetOneByUserId(userId int64) (*entity.UserEntity, error)
	AddOne(userEntity *entity.UserEntity) error
}

type UserRepository struct {
	MysqlDB *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		MysqlDB: db,
	}
}

// GetOneByUserId query userInfo by userId
func (u *UserRepository) GetOneByUserId(userId int64) (*entity.UserEntity, error) {
	userInfo := entity.UserEntity{
		UserId:     0,
		UserNick:   "",
		UserMobile: "",
		CreateTime: "",
		UpdateTime: "",
	}
	res := u.MysqlDB.Raw(GetOneByUserId, userId).Scan(&userInfo)
	return &userInfo, res.Error
}

// AddOne insert one data
func (u *UserRepository) AddOne(userEntity *entity.UserEntity) error {
	res := u.MysqlDB.Exec(AddOne, userEntity.UserNick, userEntity.UserMobile)
	return res.Error
}
