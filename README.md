# logrus 日志管理
使用"github.com/sirupsen/logrus"、"github.com/rifflock/lfshook"、"github.com/lestrrat-go/file-rotatelogs"三个库进行项目的日志管理

"github.com/rifflock/lfshook"是一个用于将日志记录到本地文件系统的Golang库。 它提供了一个钩子（hook）来扩展Golang的标准日志库，使得在日志记录时可以将日志信息写入到本地文件系统中

"github.com/sirupsen/logrus"用来记录日志

"github.com/lestrrat-go/file-rotatelogs"它提供了一个钩子（hook）来扩展标准的日志库，实现日志文件的滚动记录

## 初始化日志
logger.InitLog("log", "log", 1024*1024*100, time.Hour*24*7)

logPath：日志保存目录

logName：日志保存文件名

1024*1024*100 ： 日志保存滚动大小，这个100是100M

time.Hour*24*7：最大日志保留日期，超过这个日期则进行清理

## 使用
```
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
```