package model

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID          uint       `json:"userId" gorm:"column:user_id;"`
	Username    string     `json:"username" gorm:"type:varchar(50);comment:用户名"`
	Password    string     `json:"password" gorm:"type:varchar(200);comment:密码"`
	Nickname    string     `json:"nickname" gorm:"type:varchar(50);comment:昵称"`
	Email       string     `json:"email" gorm:"type:varchar(50);comment:邮箱"`
	AvatarUrl   string     `json:"avatarUrl" gorm:"type:varchar(100);comment:头像"`
	PhoneNumber string     `json:"phoneNumber" gorm:"type:varchar(20);comment:手机号"`
	IsFrozen    bool       `json:"isFrozen" gorm:"type:boolean;comment:是否被冻结;default:false"`
	IsAdmin     bool       `json:"isAdmin" gorm:"type:boolean;comment:是否为管理员;default:false"`
	CreateTime  *time.Time `json:"createTime"`
	UpdateTime  *time.Time `json:"updateTime"`
}

const (
	PassWordDepth = 12
)

// SetPassword 加密密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordDepth)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
