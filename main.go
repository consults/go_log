package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"test/core/logger"
	"time"
)

// 在程序初始化的时候，进行初始化一下就可以了
func main() {
	logger.InitLog("log", "log", 1024*1024*100, time.Hour*24*7)
	for {
		time.Sleep(time.Millisecond * 100)
		log.Info(fmt.Sprintf("11111111111111111111111111111111111111111111111111"))
	}
}
