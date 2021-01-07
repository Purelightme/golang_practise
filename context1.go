package main

import (
	"context"
	"fmt"
	"time"
)

func main()  {

	ctx,cancel := context.WithCancel(context.Background())

	go watch(ctx,"【监控1】")
	go watch(ctx,"【监控2】")
	go watch(ctx,"【监控3】")

	time.Sleep(time.Second * 10)

	fmt.Println("通知停止!!!")
	cancel()

	time.Sleep(time.Second * 1)

}

func watch(ctx context.Context,name string)  {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name,"监控退出...")
			return
		default:
			fmt.Println(name,"持续监控中...")
			time.Sleep(time.Second * 2)
		}
	}
}
