package main

import (
	"fmt"
	"log"
	"os"
)

func main()  {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: permission filename")
		return
	}

	stat,err := os.Stat(args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(stat.IsDir(),stat.Name(),stat.Size(),stat.Mode().Perm())
}
