package repository

import (
	"dragon/domain/entity"
	"gorm.io/gorm"
)

const (
	GetOne = `select * from t_user limit 1`
)

type IUserRepository interface {
	GetOne() (*entity.UserEntity, error)
}
type UserRepository struct {
	MysqlDB *gorm.DB
}

func NewTUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		MysqlDB: db,
	}
}

func (this *UserRepository) GetOne() (*entity.UserEntity, error) {
	var data entity.UserEntity
	res := this.MysqlDB.Raw(GetOne).Scan(&data)
	return &data, res.Error
}
