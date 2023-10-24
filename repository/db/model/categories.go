package model

import "time"

type Category struct {
	ID               uint      `json:"categoryId" gorm:"column:category_id;"`
	CategoryName     string    `json:"categoryName" gorm:"column:category_name;"`
	ParentCategoryID uint      `json:"parentCategoryId" gorm:"column:parent_category_id"`
	CategoryLevel    uint      `json:"categoryLevel" gorm:"column:category_level"`
	CreatedAt        time.Time `json:"createTime"` //字段使用CreatedAt，不是CreatedTime
	UpdatedAt        time.Time `json:"updateTime"`
}
