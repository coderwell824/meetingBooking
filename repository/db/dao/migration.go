package dao

import (
	"log"
	"meetingBooking/repository/db/model"
)

// 数据迁移
func migration() {
	//自动迁移模式
	err := _db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&model.User{},
		&model.Room{},
		&model.Booking{},
		&model.BookingAttention{})

	if err != nil {
		log.Println(err)
		return
	}
}
