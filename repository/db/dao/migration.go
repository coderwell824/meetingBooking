package dao

import (
	"log"
	"meetingBooking/repository/db/model"
)

// 数据迁移
func migration() {
	//自动迁移模式

	err := _db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(
		&model.Role{},
		&model.User{},
		&model.Permission{},
		&model.Room{},
		&model.Booking{},
		&model.Category{},
	)

	if err != nil {
		log.Println(err)
		return
	}
}
