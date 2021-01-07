package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)


func init()  {
	log.SetFormatter(&log.JSONFormatter{})
	file,_ := os.OpenFile("./test.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	log.SetOutput(file)
}

func main()  {
	log.Info("成功启动 success")

	fmt.Println("Hello,Purelight")
}