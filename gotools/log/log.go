package log

import (
	"math/rand"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugarLogger *zap.SugaredLogger

// 使用指南
/*
使用指南：
需要需要修改logPath（写死）
使用时，调用  log.Log(level, msg)
level表示信息等级
msg表示具体信息
*/

var logPath = "./test.log"
var file *os.File

func Log(level, msg string) {
	InitLogger()
	defer sugarLogger.Sync()

	switch level {
	case "Debug":
		sugarLogger.Debug(msg)
	case "Info":
		sugarLogger.Info(msg)
	case "Warn":
		sugarLogger.Warn(msg)
	case "Error":
		sugarLogger.Error(msg)
	default:
		// sugarLogger.
	}
}

func InitLogger() {
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

// 随时删除
func GetRandLevelAndMsg() (string, string) {
	levels := []string{"Debug", "Info", "Warn", "Error"}
	msgs := map[string][]string{
		"Debug": {"配置被修改了", "配置被删除", "时间自动同步", "访问了网络"},
		"Info":  {"用户进行了操作", "用户登录", "用户登出", "管理员删除了用户"},
		"Warn":  {"时间被非root用户修改了", "非管理员用户下发了扫描任务", "心跳丢失5分钟", "master节点污点被删除", "cpu超过了90%"},
		"Error": {"用户登录失败", "修改密码失败", "系统短时间内接收到上千条api请求", "系统oom", "网络不通"},
	}

	level := levels[rand.Intn(len(levels))] // 随机选择 level
	msgList := msgs[level]                  // 根据选择的 level 获取对应的 msg 列表
	msg := msgList[rand.Intn(len(msgList))] // 随机选择 msg

	return level, msg
}
