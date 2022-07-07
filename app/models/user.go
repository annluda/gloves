package models

import (
	"fmt"
	"gloves/database"
	"gloves/pkg/auth"
	"gloves/pkg/logger"
	"gloves/pkg/utils"
	"strconv"
	"time"
)

// User 用户模型
type User struct {
	BaseModel
	Name     string `gorm:"column:name;type:varchar(255);not null"`
	Email    string `gorm:"column:email;type:varchar(255);not null"`
	Avatar   string `gorm:"column:avatar;type:varchar(255);not null"`
	Password string `gorm:"column:password;type:varchar(255);not null"`
	// 是否为管理员
	IsAdmin uint `gorm:"column:is_admin;type:tinyint(1)"`
	// 用户激活
	ActivationToken string    `gorm:"column:activation_token;type:varchar(255)"`
	Activated       uint      `gorm:"column:activated;type:tinyint(1);not null"`
	EmailVerifiedAt time.Time `gorm:"column:email_verified_at"` // 激活时间
	// 用于实现记住我功能，存入 cookie 中，下次带上时，即可直接登录
	RememberToken string `gorm:"column:remember_token;type:varchar(100)"`
}

// TableName 表名
func (User) TableName() string {
	return "users"
}

// Create -
func (u *User) Create() (err error) {
	if err = u.Encrypt(); err != nil {
		logger.Warnf("用户创建失败: %v", err)
		return err
	}

	// 生成用户 remember_token
	if u.RememberToken == "" {
		u.RememberToken = string(utils.RandomCreateBytes(10))
	}
	// 生成用户激活 token
	if u.ActivationToken == "" {
		u.ActivationToken = string(utils.RandomCreateBytes(30))
	}

	if err = database.DB.Create(&u).Error; err != nil {
		logger.Warnf("用户创建失败: %v", err)
		return err
	}

	return nil
}

// Update 更新用户
func (u *User) Update(needEncryotPwd bool) (err error) {
	if needEncryotPwd {
		if err = u.Encrypt(); err != nil {
			logger.Warnf("用户更新失败: %v", err)
			return err
		}
	}

	if err = database.DB.Save(&u).Error; err != nil {
		logger.Warnf("用户更新失败: %v", err)
		return err
	}

	return nil
}

// Delete -
func Delete(id int) (err error) {
	user := &User{}
	user.BaseModel.ID = uint(id)

	// Unscoped: 永久删除而不是软删除 (由于该操作是管理员操作的，所以不使用软删除)
	if err = database.DB.Unscoped().Delete(&user).Error; err != nil {
		logger.Warnf("用户删除失败: %v", err)
		return err
	}

	return nil
}

// Encrypt 对密码进行加密
func (u *User) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

// Compare 验证用户密码
func (u *User) Compare(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)
	return
}

// UserGet Get -
func UserGet(id int) (*User, error) {
	user := &User{}
	d := database.DB.First(&user, id)
	return user, d.Error
}

// GetByEmail -
func GetByEmail(email string) (*User, error) {
	user := &User{}
	d := database.DB.Where("email = ?", email).First(&user)
	return user, d.Error
}

func GetByName(name string) (*User, error) {
	user := &User{}
	d := database.DB.Where("name = ?", name).First(&user)
	return user, d.Error
}

// GetByActivationToken -
func GetByActivationToken(token string) (*User, error) {
	user := &User{}
	d := database.DB.Where("activation_token = ?", token).First(&user)
	return user, d.Error
}

// GetByRememberToken -
func GetByRememberToken(token string) (*User, error) {
	user := &User{}
	d := database.DB.Where("remember_token = ?", token).First(&user)
	return user, d.Error
}

// List 获取用户列表
func List(offset, limit int) (users []*User, err error) {
	users = make([]*User, 0)

	if err := database.DB.Offset(offset).Limit(limit).Order("id").Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

// All -
func All() (users []*User, err error) {
	users = make([]*User, 0)

	if err := database.DB.Order("id").Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

// AllCount 总用户数
func AllCount() (count int, err error) {
	err = database.DB.Model(&User{}).Count(&count).Error
	return
}

// Gravatar 获取用户头像
func (u *User) Gravatar() string {
	return fmt.Sprintf("/public/img/%d.png", u.ID)
}

// GetIDstring 获取字符串形式的 id
func (u *User) GetIDstring() string {
	return strconv.Itoa(int(u.ID))
}

// IsAdminRole 是否为管理员
func (u *User) IsAdminRole() bool {
	return u.IsAdmin == TrueTinyint
}

// IsActivated 是否已激活
func (u *User) IsActivated() bool {
	return u.Activated == TrueTinyint
}
