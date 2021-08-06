package entity

type TUserEntity struct {
	UserId     int64 `gorm:"primaryKey;AUTO_INCREMENT"`
	UserNick   string
	UserMobile string
	Amount     float64
	CreateTime string
	UpdateTime string
}

// TableName set orm table name
func (TUserEntity) TableName() string {
	return "t_user"
}
