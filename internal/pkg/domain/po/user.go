package po

type User struct {
	ID       int    `gorm:"column:id"`
	Username string `gorm:"column:username"` // 用户名
	Password string `gorm:"column:password"` // 密码
}

func (u User) TableName() string {
	return "user"
}
