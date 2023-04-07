package main

import (
	"go-gin-example/config"
	"go-gin-example/global"
	"go-gin-example/router"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	// 设置随机数种子
	rand.Seed(time.Now().Unix())

	f := logInit()
	defer f.Close()

	// 连接数据库
	_ = global.NewDB()

	// 自动迁移
	global.AutoMigrate()

	r := router.NewRouter()
	_ = r.Run(config.AppConfig.RunPort)
}

func logInit() *os.File {
	// 打开日志文件
	f, err := os.OpenFile("./runtime/log/log.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatalln("日志文件打开失败")
	}

	// 设置日志输出到文件和控制台
	multiWriter := io.MultiWriter(f, os.Stdout)
	log.SetOutput(multiWriter)

	return f
}
