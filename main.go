package main

import (
	"fmt"
	"test/gotools/log"
	"time"
)

type Inventory struct {
	Material string
	Count    uint
}

type Test []int

func main() {
	// TdEngine的访问+操作 demo
	// td.TdSample()

	for {
		time.Sleep(1 * time.Second)

		level, msg := log.GetRandLevelAndMsg()
		log.Log(level, msg)
	}

}

func LogIn(level, msg string) {
	fmt.Printf("[%s]: msg: %s", level, msg)
}
