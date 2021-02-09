package main

import (
	"bytes"
	"fmt"
	"os"
)

func main()  {
	buf := new(bytes.Buffer)

	//buf.Write([]byte{'a','1','A'})

	//buf.Write([]byte{'我','和','你'})
	buf.WriteString("我和你")
	s := "哈皮"
	r := []rune(s)
	buf.WriteRune(r[1])

	fmt.Println(r)

	fmt.Println(buf.String())

	res,n,_ := buf.ReadRune()
	fmt.Println(string(res),n)

	buf.WriteTo(os.Stdout)
	fmt.Println(buf.String())
}
