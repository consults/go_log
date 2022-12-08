package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"test/logger"
	"time"
)

func main() {
	logger.InitLog("log", "log", time.Second*2, time.Second*6)
	for {
		time.Sleep(time.Second * 1)
		logrus.Info(fmt.Sprintf("1111"))
	}
}
