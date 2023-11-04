package model

import "time"

type Product struct {
	ID            uint      `json:"productId" gorm:"column:product_id;"`
	Name          string    `json:"productName" gorm:"size:255;index"`
	CategoryID    uint      `json:"categoryId" gorm:"not null"`
	Title         string    `json:"title" `
	Info          string    `json:"info" gorm:"size:1000;"`
	ImgPath       string    `json:"imgPath"`
	Price         float64   `json:"price"`
	DisCountPrice float64   `json:"disCountPrice"`
	OnSale        bool      `json:"onSale"` //是否在售
	Num           uint      `json:"num"`
	BossID        uint      `json:"bossId"`
	BossName      string    `json:"bossName"`
	BossAvatar    string    `json:"bossAvatar"`
	CreatedAt     time.Time `json:"createTime"` //字段使用CreatedAt，不是CreatedTime
	UpdatedAt     time.Time `json:"updateTime"`
}
