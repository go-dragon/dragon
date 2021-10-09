package service

import (
	"dragon/domain/entity"
	"dragon/domain/repository"
	"gorm.io/gorm"
)

type IUserService interface {
	GetOne() (*entity.UserEntity, error)
}
type UserService struct {
	UserRepository repository.IUserRepository
}

func NewUserService(db *gorm.DB) IUserService {
	return &UserService{
		UserRepository: repository.NewUserRepository(db),
	}
}
func (this *UserService) GetOne() (*entity.UserEntity, error) {
	var data *entity.UserEntity
	data, err := this.UserRepository.GetOne()
	return data, err
}
