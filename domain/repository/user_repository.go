package repository

import (
	"dragon/domain/entity"
	"dragon/domain/mapper/usermapper"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetOne() (*entity.UserEntity, error)
}
type UserRepository struct {
	MysqlDB *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		MysqlDB: db,
	}
}
func (this *UserRepository) GetOne() (*entity.UserEntity, error) {
	var data entity.UserEntity
	res := this.MysqlDB.Raw(usermapper.GetOne).Scan(&data)
	return &data, res.Error
}
