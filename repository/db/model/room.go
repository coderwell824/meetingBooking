package model

import "time"

type Room struct {
	ID          uint      `json:"roomId" gorm:"column:room_id;"`        //会议室Id
	Name        string    `json:"name" gorm:"type:varchar(50)"`         //会议室名称
	Capability  uint      `json:"capability"`                           //会议室容量
	Location    string    `json:"location" gorm:"type:varchar(50)"`     //会议室位置
	Equipment   string    `json:"equipment" gorm:"type:varchar(50)"`    //会议室设备
	Description string    `json:"description" gorm:"type:varchar(100)"` //会议室描述
	IsBooked    bool      `json:"isBooked"`                             //是否被预订
	CreateTime  time.Time `json:"createTime"`                           //创建时间
	UpdateTime  time.Time `json:"updateTime"`                           //更新时间
}
