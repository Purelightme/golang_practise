package main

import "fmt"

func main()  {

	minMax := func(a int,b int)(min,max int) {
		if a > b {
			min = b
			max = a
		}else {
			min = a
			max = b
		}
		return
	}

	min,max := minMax(10,3)

	fmt.Println(min,max)

	f := returnFunc()

	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

//函数
func returnFunc() func()int  {
	i := 0;

	return func() int {
		i++
		return i * i
	}
}