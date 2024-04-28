package nats

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	nats "github.com/nats-io/nats.go"
)

func Publish() {
	// 连接Nats服务器
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("无法连接到NATS服务器: %v", err)
	}
	defer nc.Close()

	res := `{"field1":"2024.04.22", "field2":"aaaa", "field3":20}`

	// 发布消息
	if err := nc.Publish("jszhou", []byte(res)); err != nil {
		log.Fatalf("无法发布消息: %v", err)
	}
	fmt.Printf("succeed to publish msg\n")
}

func Consumer() {
	// 连接Nats服务器
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("无法连接到NATS服务器: %v", err)
	}
	defer nc.Close()

	// 创建一个简单的异步订阅
	sub, err := nc.Subscribe("subject", func(msg *nats.Msg) {
		// 打印接收到的消息
		log.Printf("接收到的消息: %s", msg.Data)
	})
	if err != nil {
		log.Fatalf("无法订阅: %v", err)
	}

	// 创建一个channel来接收退出信号
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// 进入循环，等待退出信号
	for {
		select {
		case <-done:
			// 收到退出信号，取消订阅并退出循环
			sub.Unsubscribe()
			fmt.Println("程序退出")
			return
		}
	}
}
