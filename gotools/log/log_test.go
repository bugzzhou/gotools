package log

import (
	"math/rand"
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	for {
		time.Sleep(1 * time.Second)

		level, msg := getRandLevelAndMsg()
		Log(level, msg)
	}
}

// 随时删除
func getRandLevelAndMsg() (int, string) {
	levels := []int{Debug, Info, Warn, Error}
	msgs := map[int][]string{
		Debug: {"配置被修改了", "配置被删除", "时间自动同步", "访问了网络"},
		Info:  {"用户进行了操作", "用户登录", "用户登出", "管理员删除了用户"},
		Warn:  {"时间被非root用户修改了", "非管理员用户下发了扫描任务", "心跳丢失5分钟", "master节点污点被删除", "cpu超过了90%"},
		Error: {"用户登录失败", "修改密码失败", "系统短时间内接收到上千条api请求", "系统oom", "网络不通"},
	}

	level := levels[rand.Intn(len(levels))] // 随机选择 level
	msgList := msgs[level]                  // 根据选择的 level 获取对应的 msg 列表
	msg := msgList[rand.Intn(len(msgList))] // 随机选择 msg

	return level, msg
}
