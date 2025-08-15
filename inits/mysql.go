package inits

import (
	"fmt"
	"github.com/yuanzhichao-star/public_pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"

	"sync"
	"time"
)

func InitMysql() {
	var err error
	var Once sync.Once //设置单例模式
	mysqlCong := config.AppCong.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlCong.User, mysqlCong.Password, mysqlCong.Host, mysqlCong.Port, mysqlCong.Database)

	//单例模式
	Once.Do(func() {
		config.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("mysql init fail")
		}
		log.Println("mysql init success")
	})

	//设置连接池
	sqlDB, err := config.DB.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
}
