package main

import (
	"fmt"
	"log"
	"net"
)

func main()  {

	i,err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}

	for _,item := range i{
		fmt.Println(item.Name)
		fmt.Println(item.Flags.String())
		fmt.Println(item.MTU)
		fmt.Println(item.HardwareAddr.String())
		byName,err := net.InterfaceByName(item.Name)
		if err != nil {
			log.Fatal(err)
		}

		addrs,err := byName.Addrs()
		if err != nil {
			log.Fatal(err)
		}

		for index,addr := range addrs {
			fmt.Println("Addr #",index,addr.String())
		}

		fmt.Println("============================")
	}

}
