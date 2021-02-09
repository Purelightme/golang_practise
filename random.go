package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	rand.Seed(time.Now().Unix())
	fmt.Println(rand.Intn(100))

	start := "!"
	fmt.Println(byte(start[0]))
	fmt.Println(string(start[0] + byte(rand.Intn(94))))
}
