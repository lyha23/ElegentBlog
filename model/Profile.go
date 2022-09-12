package model

type Profile struct {
	User
	Name    string `gorm:"type:varchar(20)" json:"name"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"`
	Qqchat  string `gorm:"type:varchar(200)" json:"qq_chat"`
	Wechat  string `gorm:"type:varchar(100)" json:"wechat"`
	Bili    string `gorm:"type:varchar(200)" json:"bili"`
	Github  string `gorm:"type:varchar(200)" json:"github"`
	Email   string `gorm:"type:varchar(200)" json:"email"`
	Avatar  string `gorm:"type:varchar(200)" json:"avatar"`
	AboutMe string `gorm:"type:varchar(3000)" json:"aboutMe"`
}

// GetProfile 获取个人信息设置
func (ctl *Profile) GetProfile() error {
	err = db.Where("ID = ?", ctl.ID).Omit("password").FirstOrCreate(&ctl).Error
	return err
}

// UpdateProfile 更新个人信息设置
func (ctl *Profile) UpdateProfile() error {
	profile := *ctl
	err = db.Where("ID = ?", ctl.ID).FirstOrCreate(ctl).Error
	err = db.Model(&ctl).Where("ID = ?", ctl.ID).Updates(profile).Error
	return err
}