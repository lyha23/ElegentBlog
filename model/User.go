package model

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password,omitempty" validate:"required,min=6,max=120" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role,omitempty" validate:"required,gte=2" label:"角色码"`
}

type UserEvent struct {
	User     `json:"user"`
	UserList []User `json:"userList"`
}

func (ctl *User) AddUserRecordIfNeeded() error {
	err := db.Table("user").Where("username=?", ctl.Username).First(&ctl).Error
	if ctl.ID == 0 {
		err = db.Table("user").Create(&ctl).Error
		return err
	}
	return err
}

func (ctl *User) SaveUser() error {
	if ctl.ID != 0 {
		err = db.Model(&ctl).Where("id=?", ctl.ID).Updates(ctl).Error
	} else {
		_ = db.Model(&ctl).Where("username=?", ctl.Username).First(&ctl).Error
		if ctl.ID == 0 {
			err = db.Create(&ctl).Error
		} else {
			err = fmt.Errorf("用户重名")
		}
	}
	return err
}

// CheckUser 查询用户是否存在
func (ctl *User) CheckUser() error {
	err := db.Select("id").Where("username = ?", ctl.Username).First(&ctl).Error
	return err
}

// CreateUser 新增用户
func (ctl *User) CreateUser() error {
	err := db.Create(&ctl).Error
	return err
}

// GetUser 查询用户
func (ctl *User) GetUserByID() error {
	err := db.Limit(1).Where("ID = ?", ctl.ID).Find(&ctl).Error
	return err
}

// GetUsers 查询用户列表
func (ctl *UserEvent) GetUsers() error {
	if ctl.Username != "" {
		err = db.Select("id,username,role,created_at").Where(
			"username LIKE ?", ctl.Username+"%",
		).Find(&ctl.UserList).Error
	} else {
		err = db.Select("id,username,role,created_at").Find(&ctl.UserList).Error
	}
	return err
}

// EditUser 编辑用户信息
func (ctl *User) EditUser() error {
	var maps = make(map[string]interface{})
	maps["username"] = ctl.Username
	maps["role"] = ctl.Role
	err = db.Model(&ctl).Where("id = ? ", ctl.ID).Updates(maps).Error
	return err
}

// ChangePassword 修改密码
func (ctl *User) ChangePassword() error {
	err = db.Select("password").Where("id = ?", ctl.ID).Updates(&ctl).Error
	return err
}

// DeleteUser 删除用户
func (ctl *User) DeleteUser() error {
	var user User
	err = db.Where("id = ? ", ctl.ID).Delete(&user).Error
	return err
}

func (ctl *User) BeforeUpdate(_ *gorm.DB) error {
	err := ctl.ScryptPw()
	return err
}

// BeforeCreate 密码加密&权限控制
func (ctl *User) BeforeCreate(*gorm.DB) error {
	err := ctl.ScryptPw()
	ctl.Role = 2
	return err
}

// ScryptPw 生成密码
func (ctl *User) ScryptPw() error {
	const cost = 10
	HashPw, err := bcrypt.GenerateFromPassword([]byte(ctl.Password), cost)
	ctl.Password = string(HashPw)
	return err
}

// CheckLogin 后台登录验证
func (ctl *User) CheckLogin() error {
	user := *ctl
	err = db.Where("username = ?", ctl.Username).First(&ctl).Error
	err = bcrypt.CompareHashAndPassword([]byte(ctl.Password), []byte(user.Password))
	return err
}

// CheckLoginFront 前台登录
func (ctl *User) CheckLoginFront() error {
	var user User
	db.Where("username = ?", ctl.Username).First(&user)
	err = bcrypt.CompareHashAndPassword([]byte(ctl.Password), []byte(user.Password))
	if err == nil {
		db.Where("username = ?", ctl.Username).First(&ctl)
	}
	return err
}