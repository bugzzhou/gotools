package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/*
	使用指南：

init中修改文件路径
log.Log(level, msg) 直接用于存写日志
log.GetWriteLevel() 用于获取当前记录等级
log.ChangeWriteLevel(level) 用于修改记录等级
*/

const (
	Debug = iota
	Info
	Warn
	Error
)

var sugarLogger *zap.SugaredLogger
var defaultLogPath = "./test.log"
var file *os.File
var writeLevel int

var levelNames = map[int]string{
	Debug: "Debug",
	Info:  "Info",
	Warn:  "Warn",
	Error: "Error",
}

// 初始化时，需要指定日志的文件位置
func init() {
	InitLogger("", 0)
}

func Log(level int, msg string) {
	defer sugarLogger.Sync()

	if writeLevel > level {
		return
	}

	switch level {
	case Debug:
		sugarLogger.Debug(msg)
	case Info:
		sugarLogger.Info(msg)
	case Warn:
		sugarLogger.Warn(msg)
	case Error:
		sugarLogger.Error(msg)
	default:

	}
}

func GetWriteLevel() (int, string) {
	return writeLevel, levelNames[writeLevel]
}

func ChangeWriteLevel(level int) {
	writeLevel = level
}

func InitLogger(logPath string, level int) {
	if level == 0 {
		writeLevel = Info
	} else {
		writeLevel = level
	}

	if logPath == "" {
		logPath = defaultLogPath
	}

	if file == nil {
		var err error
		file, err = os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
	}

	core := zapcore.NewCore(
		getEncoder(),
		getLogWriter(),
		zapcore.DebugLevel)

	logger := zap.New(core)
	sugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	return zapcore.AddSync(file)
}
