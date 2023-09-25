package dao

import (
	"log"
	"meetingBooking/repository/db/model"
)

// 数据迁移
func migration() {
	//自动迁移模式
	_db.SetupJoinTable(&model.Role{}, "Users", &model.UserRole{})
	err := _db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(
		&model.Role{},
		&model.User{},
		&model.UserRole{},
		&model.Room{},
		&model.Booking{},
	)
	
	if err != nil {
		log.Println(err)
		return
	}
}
