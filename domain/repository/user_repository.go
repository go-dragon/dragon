package repository

import (
	"dragon/domain/entity"
	"gorm.io/gorm"
)

// IUserRepository 用户仓库
type IUserRepository interface {
	IBaseRepository
}

type UserRepository struct {
	BaseRepository
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		BaseRepository: BaseRepository{
			TableName: entity.UserEntity{}.TableName(),
			MysqlDB:   db,
		},
	}
}
