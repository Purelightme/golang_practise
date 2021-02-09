package main

import (
	"fmt"
	"log"
	"time"
)

func main()  {
	t,err := time.Parse("15:04","09:02")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t.Format("2016-01-02 15:04:05"))
}
