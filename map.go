package main

import (
	"fmt"
	"time"
)

func main()  {

	user := make(map[string]interface{})

	for i := 0;i < 10;i++ {
		go doMap(user,i)
	}

	time.Sleep(time.Second)

	fmt.Println(user)
}

func doMap(u map[string]interface{},i int) {
	u["name"] = i
}


