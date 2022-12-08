package logger

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type LogFileFormatter struct {
}

func (l LogFileFormatter) Format(entry *log.Entry) ([]byte, error) {
	funcs := strings.SplitN(entry.Caller.Function, ".", 2)
	msg := fmt.Sprintf("[%s] [%s] [%s:%d] %s\n",
		entry.Time.Format("2006-01-02 15:04:05,000"),
		strings.ToUpper(entry.Level.String()),
		funcs[0]+"/"+filepath.Base(entry.Caller.File),
		entry.Caller.Line,
		entry.Message)
	return []byte(msg), nil
}
func InitLog(logPath string, logName string, rotatime time.Duration, maxage time.Duration) {
	filepaths := path.Join(logPath, logName)
	writer, _ := rotatelogs.New(
		filepaths+".%Y%m%d%H%M",               //分割后文件名
		rotatelogs.WithLinkName(filepaths),    //软连接指向最新文件
		rotatelogs.WithMaxAge(maxage),         // 最大保留时间
		rotatelogs.WithRotationTime(rotatime), // 设置切割时间
	)
	writeMap := lfshook.WriterMap{
		log.TraceLevel: writer,
		log.DebugLevel: writer,
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}
	log.AddHook(lfshook.NewHook(writeMap, new(LogFileFormatter)))

	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)
	log.SetFormatter(&LogFileFormatter{})
}
