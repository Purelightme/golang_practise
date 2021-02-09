package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

var envWg sync.WaitGroup

func main()  {

	envWg.Add(10000)

	for i := 0;i < 1000 ;i++  {
		go Start(10)
	}

	envWg.Wait()

	fmt.Println("执行完成")
}

func Start(n int)  {
	for i := 0;i < n;i++ {
		res, err := http.Get("https://test.daishutijian.com/services_component/api/frontend/common/wechatConfig?url=https%3A%2F%2Ftest.daishutijian.com%2Fservices_component%2Fhealth%2F")
		if err != nil {
			fmt.Println("Fatal error ", err.Error())
		}
		defer res.Body.Close()

		content, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Fatal error ", err.Error())
		}

		fmt.Println(string(content))
		envWg.Done()
	}
}