package model

import (
	"gorm.io/gorm"
	"math"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid" json:"category,omitempty"`
	gorm.Model
	Title        string `gorm:"type:varchar(100);not null" json:"title"`
	Cid          int    `gorm:"type:int;not null" json:"cid"`
	Desc         string `gorm:"type:varchar(200)" json:"desc"`
	Content      string `gorm:"type:longtext" json:"content"`
	Img          string `gorm:"type:varchar(100)" json:"img"`
	CommentCount int    `gorm:"type:int;not null;default:0" json:"commentCount,omitempty"`
	ReadCount    int    `gorm:"type:int;not null;default:0" json:"readCount,omitempty"`
}

type Pagination struct {
	Page      int   `json:"page"`      // 页码
	PageSize  int   `json:"pageSize"`  // 每页大小
	Total     int64 `json:"total"`     // 总记录数
	TotalPage int   `json:"totalPage"` // 总页数
}

type Event struct {
	Article     `json:"article"`
	Pagination  `json:"pagination"`
	ArticleList []Article `json:"articleList"`
}

func (p *Pagination) DefaultRequest() {
	switch {
	case p.PageSize >= 100:
		p.PageSize = 100
	case p.PageSize <= 0:
		p.PageSize = 10
	}
	if p.Page == 0 {
		p.Page = 1
	}
}

// SaveArt 保存文章
func (ctl *Article) SaveArt() error {
	if ctl.ID == 0 {
		err = db.Create(&ctl).Error
	} else {
		var maps = make(map[string]interface{})
		maps["title"] = ctl.Title
		maps["cid"] = ctl.Cid
		maps["desc"] = ctl.Desc
		maps["content"] = ctl.Content
		maps["img"] = ctl.Img
		err = db.Model(&ctl).Where("id = ? ", ctl.ID).Updates(&maps).Error
	}
	return err
}

// GetCateArt 查询分类下的所有文章
func (ctl *Event) GetCateArt() error {
	err = db.Preload("Category").Limit(ctl.PageSize).Offset((ctl.Page-1)*ctl.PageSize).Where(
		"cid =?", ctl.Cid).Find(&ctl.ArticleList).Error
	db.Model(&ctl.ArticleList).Where("cid =?", ctl.Cid).Count(&ctl.Total)
	return err
}

// GetArtInfo 查询单个文章
func (ctl *Article) GetArtInfo() error {
	err = db.Where("id = ?", ctl.ID).Preload("Category").First(&ctl).Error
	db.Model(&ctl).Where("id = ?", ctl.ID).UpdateColumn("read_count", gorm.Expr("read_count + ?", 1))
	return err
}

// GetArt 查询文章列表
func (ctl *Event) GetList() error {
	if len(ctl.Title) > 0 {
		err = db.Select("article.id,title, img, created_at, updated_at, `desc`, comment_count, read_count, Category.name").Order("id desc").Joins("Category").Where("title LIKE ?",
			ctl.Title+"%",
		).Limit(ctl.PageSize).Offset((ctl.Page - 1) * ctl.PageSize).Find(&ctl.ArticleList).Error
		db.Model(&ctl.ArticleList).Where("title LIKE ?",
			ctl.Title+"%",
		).Count(&ctl.Total)
	} else {
		err = db.Select("article.id, title, img, content,created_at, updated_at, `desc`, comment_count, read_count, Category.name").Limit(ctl.PageSize).Offset((ctl.Page - 1) * ctl.PageSize).Order("id DESC").Joins("Category").Find(&ctl.ArticleList).Error
		// 单独计数
		db.Model(&ctl.ArticleList).Count(&ctl.Total)
	}
	ctl.TotalPage = int(math.Ceil(float64(ctl.Total) / float64(ctl.PageSize)))
	return err
}

func (ctl *Event) GetArtTimeline() error {
	err := db.Select("article.id,title,created_at").Find(&ctl.ArticleList).Error
	return err
}

// DeleteArt 删除文章
func (ctl *Article) DeleteArt() error {
	err = db.Where("id = ? ", ctl.ID).Delete(&ctl).Error
	return err
}