package main

import (
	"fmt"
	"time"
)

func main()  {

	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("监控退出...")
				return
			default:
				fmt.Println("持续监控中...")
				time.Sleep(time.Second * 2)
			}
		}
	}()

	time.Sleep(time.Second * 10)
	fmt.Println("通知停止!!!")
	stop <- true

	time.Sleep(time.Second * 4)

}