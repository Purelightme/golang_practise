package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var benefitWg sync.WaitGroup

func main()  {
	fmt.Println(time.Now())
	benefitWg.Add(1000)
	go StartRequest(100,1)
	go StartRequest(100,2)
	go StartRequest(100,3)
	go StartRequest(100,4)
	go StartRequest(100,5)
	go StartRequest(100,6)
	go StartRequest(100,7)
	go StartRequest(100,8)
	go StartRequest(100,9)
	go StartRequest(100,10)
	benefitWg.Wait()
	fmt.Println(time.Now())
}

func StartRequest(n int,seq int)  {
	for i := 0;i < n;i++ {
		body := "{\"kangaroo_id\": \"oraja0zK6a74ipTHKOYg9M4VcRxA\",\"amount\": "+strconv.Itoa(10)+",\"allocateNo\":\""+strconv.Itoa(seq)+"ioi6542s53t8r6trofees34254"+strconv.Itoa(i)+"\""+",\"ts\": 1610590975,\"nonce\": \"BtKGWLTV\",\"sign\": \"47f977b99be7138c1f606b59c3a234b148898172\",\"src\": \"1\",\"reason\":\"uu\"}"
		res, err := http.Post("http://test.daishutijian.com/services_component/not-open-platform/benefit/add_point", "application/json;charset=utf-8", bytes.NewBuffer([]byte(body)))
		if err != nil {
			fmt.Println("Fatal error ", err.Error())
		}
		defer res.Body.Close()

		content, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Fatal error ", err.Error())
		}

		fmt.Println(string(content))
		benefitWg.Done()
	}
}