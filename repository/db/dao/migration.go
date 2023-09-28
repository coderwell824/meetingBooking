package dao

import (
	"log"
	"meetingBooking/repository/db/model"
)

// 数据迁移
func migration() {
	//自动迁移模式
	_db.SetupJoinTable(&model.Role{}, "Permissions", &model.RolePermission{})
	_db.SetupJoinTable(&model.Permission{}, "Roles", &model.RolePermission{})
	//_db.SetupJoinTable(&model.User{}, "Bookings", &model.Booking{})
	//_db.SetupJoinTable(&model.Booking{}, "Users", &model.User{})
	_db.SetupJoinTable(&model.Role{}, "Users", &model.User{})

	err := _db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(
		&model.Role{},
		&model.User{},
		&model.Permission{},
		&model.Room{},
		&model.Booking{},
	)

	if err != nil {
		log.Println(err)
		return
	}
}
