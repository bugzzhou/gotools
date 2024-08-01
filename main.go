package main

import (
	"fmt"
)

type Inventory struct {
	Material string
	Count    uint
}

type Test []int

func main() {
	// TdEngine的访问+操作 demo
	// td.TdSample()

}

func LogIn(level, msg string) {
	fmt.Printf("[%s]: msg: %s", level, msg)
}
