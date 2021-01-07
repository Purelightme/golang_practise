package main

import (
	"fmt"
	"sync"
)

var wg1 sync.WaitGroup
var user sync.Map

func main()  {

	wg1.Add(10)
	for i := 0;i< 10;i++ {
		go doSyncMap(i)
	}

	wg1.Wait()

	user.Range(func(key, value interface{}) bool {
		fmt.Println(key," : ",value)
		return true
	})
}

func doSyncMap(i int)  {
	user.Store("age",i)
	v,_ := user.Load("age")
	fmt.Println("age is ",v)
	wg1.Done()
}
