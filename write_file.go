package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func main()  {

	dir,err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	f1,err := os.Create(path.Join(dir,"test_files","f1.txt"))
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()
	fmt.Fprintf(f1,"1+1=%s","2")

	f2,err := os.Create(path.Join(dir,"test_files","f2.txt"))
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()
	n,err := f2.WriteString("哈哈")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("f2写入字节数：",n)

	f3,err := os.Create(path.Join(dir,"test_files","f3.txt"))
	if err != nil {
		log.Fatal(err)
	}
	defer f3.Close()
	writer := bufio.NewWriter(f3)
	n,err = writer.WriteString("f3哈哈")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("f3写入字节数：",n)

	f4 := path.Join(dir,"test_files","f4.txt")
	err = ioutil.WriteFile(f4,[]byte{'A','a'},0644)
	if err != nil {
		log.Fatal(err)
	}

	f5,err := os.Create(path.Join(dir,"test_files","f5.txt"))
	if err != nil {
		log.Fatal(err)
	}
	defer f5.Close()
	n,err = io.WriteString(f5,"f5---io")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("f5写入字节数：",n)
}
