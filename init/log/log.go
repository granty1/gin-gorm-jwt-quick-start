package log

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-cli/tools"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// log 输出初始化
// 控制log的输出方式，输出位置，输出格式

// 用到的核心包
// logrus  		https://www.github.com/sirupsen/logrus
// rotatelogs 	https://github.com/lestrrat/go-file-rotatelogs
// lfshook 		https://github.com/rifflock/lfshook

//Log 公共Log对象，对外暴露
var Log = logrus.New()

//Init 初始化logrus
func Init() {
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("log file open fail", err)
	}
	Log.SetOutput(src)
	Log.SetLevel(logrus.DebugLevel)
	if ok := tools.PathIsExist("./log"); !ok {
		fmt.Println("create log directory")
		os.Mkdir("log", os.ModePerm)
	}
	logFilePrefix := "./log/api.log"
	logWriter, err := rotatelogs.New(
		logFilePrefix+".%Y-%m-%d-%H-%M.log",
		rotatelogs.WithLinkName(logFilePrefix),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	writerMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.ErrorLevel: logWriter,
	}
	hook := lfshook.NewHook(writerMap, &logrus.JSONFormatter{})
	Log.AddHook(hook)
}
