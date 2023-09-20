package dao

import (
	"log"
)

// 数据迁移
func migration() {
	//自动迁移模式
	err := _db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate()
	
	if err != nil {
		log.Println(err)
		return
	}
}
