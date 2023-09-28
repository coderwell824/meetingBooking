package dao

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"meetingBooking/config"
	"strings"
	"time"
)

var _db *gorm.DB

func MySqlInit() {
	connection := strings.Join([]string{config.DbUser, ":", config.DbPassword, "@tcp(", config.DbHost, ":", config.DbPort, ")/", config.DbName, "?charset=utf8&parseTime=True&loc=Local"}, "")
	//数据库日志
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connection,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: ormLogger, //打印日志
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表明不加s
		},
	})

	if err != nil {
		panic(err)
	}

	sqlDb, _ := db.DB()
	sqlDb.SetMaxOpenConns(20)  // 设置连接池，空闲
	sqlDb.SetMaxOpenConns(100) // 打开
	sqlDb.SetConnMaxLifetime(30 * time.Second)
	_db = db
	fmt.Println("mysql连接成功")
	migration()

}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
