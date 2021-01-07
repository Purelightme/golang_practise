package main

import (
	"fmt"
	"sync"
	"time"
)

var rwLock sync.RWMutex

func main()  {
	for i := 0;i < 5 ;i++  {
		go func(n int) {
			rwLock.RLock()
			defer rwLock.RUnlock()
			fmt.Printf("%d 获取到读锁：%s\n",n,time.Now().String())
			time.Sleep(time.Second)
		}(i)
	}

	time.Sleep(time.Second/10) //保证前5个go routine先执行

	for j := 0;j < 5;j++ {
		go func(n int) {
			rwLock.Lock()
			defer rwLock.Unlock()
			fmt.Printf("%d 获取到写锁：%s\n",n,time.Now().String())
			time.Sleep(time.Second)
		}(j)
	}

	time.Sleep(time.Second * 10)
}
