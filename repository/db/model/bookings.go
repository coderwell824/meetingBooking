package model

import "time"

type Booking struct {
	ID         uint      `json:"booksId" gorm:"column:book_id;"` //预定ID
	UserID     uint      `json:"userId"`                         //预定用户ID
	RoomID     uint      `json:"roomId"`                         //会议室ID
	StartTime  time.Time `json:"startTime"`                      //会议开始时间
	EndTime    time.Time `json:"endTime"`                        //会议结束时间
	Status     string    `json:"status"`                         //状态 0：申请中 1：审批通过 2：审批驳回 3：已解除
	Note       string    `json:"note" gorm:"type:varchar(100)"`  //备注
	CreateTime time.Time `json:"createTime"`                     //创建时间
	UpdateTime time.Time `json:"updateTime"`                     //更新时间
}

//BookingAttention 预定-参会中间表
type BookingAttention struct {
	ID        uint `json:"id"`
	UserID    uint `json:"userId"`
	BookingID uint `json:"bookingId"`
}
