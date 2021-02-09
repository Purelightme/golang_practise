package main

import (
	"bufio"
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main()  {
	flag.Parse()

	for _,file := range flag.Args()  {
		err := lineByLine(file)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func lineByLine(file string) error  {
	var err error;
	
	f,err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	
	r := bufio.NewReader(f)

	var seed int64
	err = binary.Read(r,binary.BigEndian,&seed)
	if err != nil {
		return err
	}
	fmt.Println(seed)

	for{
		line,err := r.ReadString('\n')

		if err == io.EOF {
			break
		}else if err != nil{
			return err
		}
		//fmt.Print(line)
		for _,x := range line  {
			fmt.Print(string(x))
		}
	}
	return nil
}