package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main()  {
	wg.Add(5)

	for i := 0;i < 5;i++ {
		go func(n int) {
			fmt.Printf("%d at %s \n",n,time.Now().String())
			//time.Sleep(time.Second)
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Println("终于结束了"+time.Now().String())
}
