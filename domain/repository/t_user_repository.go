package repository

import (
	"dragon/domain/entity"
	"gorm.io/gorm"
)

const (
	GetOne = `select * from t_user limit 1`
)

type ITUserRepository interface {
	GetOne() (*entity.TUserEntity, error)
}
type TUserRepository struct {
	MysqlDB *gorm.DB
}

func NewTUserRepository(db *gorm.DB) ITUserRepository {
	return &TUserRepository{
		MysqlDB: db,
	}
}
func (this *TUserRepository) GetOne() (*entity.TUserEntity, error) {
	var data entity.TUserEntity
	res := this.MysqlDB.Raw(GetOne).Scan(&data)
	return &data, res.Error
}
