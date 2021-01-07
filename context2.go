package main

import (
	"context"
	"fmt"
	"time"
)

func main()  {
	
	ctx,cancel := context.WithTimeout(context.Background(),time.Second * 3)
	
	defer cancel()

	go handle(ctx,time.Millisecond * 500)

	select {
	case <- ctx.Done():
		fmt.Println("结束")
	}
	
}

func handle(ctx context.Context,duration time.Duration)  {
	select {
	case <- ctx.Done():
		fmt.Println(ctx.Err())
	case <- time.After(duration):
		fmt.Println("request 处理时长 ",duration)
	}
}
