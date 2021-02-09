package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main()  {
	str := "\x99\x42\x32"
	fmt.Println(str)
	fmt.Printf("%x\n",str)
	fmt.Printf("len:%d\n",len(str))
	fmt.Printf("%q\n",str)
	fmt.Printf("%+q\n",str)
	fmt.Printf("% x\n",str)
	fmt.Printf("%s\n",str)
	fmt.Printf("%b\n",7)
	fmt.Printf("%c\n",37)
	fmt.Printf("%x\n", "abc")
	fmt.Printf("%q\n","string")
	p := Point{1, 2}
	fmt.Printf("%v\n",p)
	fmt.Printf("%+v\n",p)
	fmt.Printf("%#v\n",p)
	fmt.Printf("%T\n",p)
	fmt.Printf("%p\n", &p)
	fmt.Printf("|%5d|%5d|\n", 12, 345)
	fmt.Printf("|%5.2f|%5.2f|\n", 1.2, 3.45)
	fmt.Printf("|%-5.2f|%-5.2f|\n", 1.2, 3.45)

}
