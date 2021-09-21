package service

import (
	"dragon/domain/entity"
	"dragon/domain/repository"
	"gorm.io/gorm"
)

type ITUserService interface {
	GetOne() (*entity.UserEntity, error)
}
type TUserService struct {
	TUserRepository repository.ITUserRepository
}

func NewTUserService(db *gorm.DB) ITUserService {
	return &TUserService{
		TUserRepository: repository.NewTUserRepository(db),
	}
}
func (this *TUserService) GetOne() (*entity.UserEntity, error) {
	var data *entity.UserEntity
	data, err := this.TUserRepository.GetOne()
	return data, err
}
