package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(time.Second)
	afterCh := time.After(5 * time.Second)

	timeoutFlag := false

	for !timeoutFlag {
		select {
		case <-ticker.C:
			fmt.Println("1秒过去了")
		case <-afterCh:
			fmt.Println("倒计时三秒结束,程序退出")
			timeoutFlag = true
		}
	}

	fmt.Println("程序结束")

}
