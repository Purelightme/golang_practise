package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main()  {
	sigs := make(chan os.Signal,1)
	signal.Notify(sigs)

	go func() {
		for {
			s := <- sigs
			switch s {
			case os.Interrupt:
				fmt.Println("退出")
				os.Exit(0)
			default:
				handleSignal(s)
			}
		}
	}()

	time.Sleep(time.Second * 60)
}

func handleSignal(signal os.Signal)  {
	fmt.Println(signal)
}
