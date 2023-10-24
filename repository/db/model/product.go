package model

import "time"

type Product struct {
	ID         uint   `json:"productId" gorm:"column:product_id;"`
	Name       string `json:"productName" gorm:"size:255;index"`
	CategoryID uint   `json:"categoryId" gorm:"not null"`
	Title      string `json:"title" `

	CreatedAt time.Time `json:"createTime"` //字段使用CreatedAt，不是CreatedTime
	UpdatedAt time.Time `json:"updateTime"`
}
