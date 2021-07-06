package entity

const (
	UserStatusDelete = 0
	UserStatusOK     = 1
)

type UserEntity struct {
	UserId     int64 `gorm:"primaryKey;AUTO_INCREMENT"`
	UserNick   string
	UserMobile string
	CreateTime string
	UpdateTime string
}

// TableName set orm table name
func (UserEntity) TableName() string {
	return "t_user"
}
