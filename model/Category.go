package model

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id,omitempty"`
	Name string `gorm:"type:varchar(20);not null" json:"name,omitempty"`
}

type CatEvent struct {
	Category
	CategoryList []Category `json:"catList"`
	Total        int64      `json:"total"`
}

// CheckCategory 查询分类是否存在
func (ctl *Category) CheckCategory() error {
	err := db.Select("id").Where("name = ?", ctl.Name).First(&ctl).Error
	return err
}

// CreateCate 新增分类
func (ctl *Category) SaveCate() error {
	if ctl.ID == 0 {
		err = db.Create(&ctl).Error
	} else {
		var maps = make(map[string]interface{})
		maps["name"] = ctl.Name
		err = db.Model(&ctl).Where("id = ? ", ctl.ID).Updates(maps).Error
	}
	return err
}

// GetCate 查询分类列表
func (ctl *CatEvent) GetList() error {
	err = db.Find(&ctl.CategoryList).Error
	db.Model(&ctl.CategoryList).Count(&ctl.Total)
	return err
}

// DeleteCate 删除分类
func (ctl *Category) DeleteCate() error {
	err = db.Where("id = ? ", ctl.ID).Delete(&ctl).Error
	return err
}