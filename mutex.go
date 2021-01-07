package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex


func main()  {
	for  {
		go run()
		time.Sleep(time.Second/5)
	}
}

func run()  {
	lock.Lock()
	defer lock.Unlock()
	fmt.Println("获取到锁："+time.Now().String())
	time.Sleep(time.Second)
	fmt.Println("释放锁："+time.Now().String())
}