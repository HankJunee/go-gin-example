package global

import (
	"fmt"
	"go-gin-example/config"
	"go-gin-example/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"net/url"
	"strings"
)

var (
	DB *gorm.DB
)

func NewDB() *gorm.DB {
	// 连接数据库. 如果数据库不存在，则创建数据库
	db, err := newDB(createDSN(config.DBConfig.DBName))

	if err != nil {
		log.Fatalf("数据库连接失败: %s\n", err.Error())
	}

	sqlDB, _ := db.DB()

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	DB = db
	return db
}

func AutoMigrate() {
	_ = DB.AutoMigrate(&model.User{})
}

func newDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: config.DBConfig.TablePrefix,
		},
	})

	if err != nil && strings.Contains(err.Error(), "Unknown database") {
		log.Println("数据库不存在，尝试创建数据库")
		tempDB, err := newDB(createDSN("test"))

		if err != nil {
			return nil, err
		}

		err = tempDB.Exec("CREATE DATABASE " + config.DBConfig.DBName).Error

		if err != nil {
			return nil, err
		}

		log.Printf("数据库 %s 创建成功\n", config.DBConfig.DBName)

		return newDB(createDSN(config.DBConfig.DBName))
	}

	return db, err
}

func createDSN(dbName string) string {
	c := config.DBConfig
	timeZone := url.QueryEscape(c.TimeZone)

	if dbName == "" {
		dbName = c.DBName
	}

	// 连接数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=%s", c.UserName, c.Password, c.DbHost, c.DbPort, dbName, timeZone)

	return dsn
}
