package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main()  {

	args := os.Args

	if len(args) != 3 {
		fmt.Println("<buf size> <file name>")
		return
	}

	size,err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal(err)
	}

	f,err := os.Open(args[2])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for {
		buf := readSize(f,size)
		if buf != nil {
			fmt.Println(string(buf))
		}else {
			break
		}
	}
}

func readSize(f *os.File,size int) []byte  {
	buf := make([]byte,size)

	n,err := f.Read(buf)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("读到的大小：",n)

	return buf[0:n]

}
