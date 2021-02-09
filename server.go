package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main()  {
	fmt.Println(os.Getpid())

	listen,err := net.Listen("tcp","0.0.0.0:9009")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn,err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go func() {
			defer conn.Close()

			for {
				buf := make([]byte,1)
				n,err := conn.Read(buf)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("读取到数据：",string(buf[:n]))
			}
		}()
	}
}
