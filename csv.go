package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main()  {
	f,err := os.Create("demo.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//f.WriteString("xEFxBBxBF")
	writer := csv.NewWriter(f)
	writer.Write([]string{"编号","姓名","年龄"})
	writer.Write([]string{"1","张三","20"})
	writer.Write([]string{"2","李四","21"})
	writer.Flush()
}