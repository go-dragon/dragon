package service

import (
	"dragon/domain/entity"
	"dragon/domain/repository"
	"gorm.io/gorm"
)

type ITUserService interface {
	GetOne() (*entity.TUserEntity, error)
}
type TUserService struct {
	TUserRepository repository.ITUserRepository
}

func NewTUserService(db *gorm.DB) ITUserService {
	return &TUserService{
		TUserRepository: repository.NewTUserRepository(db),
	}
}
func (this *TUserService) GetOne() (*entity.TUserEntity, error) {
	var data *entity.TUserEntity
	data, err := this.TUserRepository.GetOne()
	return data, err
}
