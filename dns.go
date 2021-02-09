package main

import (
	"fmt"
	"log"
	"net"
)

func main()  {
	ip := "127.0.0.1"
	host := "dev.daishutijian.com"
	IP := net.ParseIP(ip)
	ips,err := lookIp(IP.String())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ips)

	hosts,err := lookHostName(host)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hosts)
}

func lookIp(ip string)([]string,error)  {
	ips,err := net.LookupAddr(ip)
	if err != nil {
		log.Fatal(err)
	}
	return ips,nil
}

func lookHostName(host string)([]string,error)  {
	hosts,err := net.LookupHost(host)
	if err != nil {
		return nil,err
	}
	return hosts,nil
}
