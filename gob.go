package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

var DATA_FILE = "test_files/data.txt"

func main()  {
	//save()
	load()
}

func save()  {
	f,err := os.Create(DATA_FILE)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	encoder := gob.NewEncoder(f)
	err = encoder.Encode("嘟嘟,{\"id\":1}")
	if err != nil {
		log.Fatal(err)
	}
}

func load()  {
	f,err := os.Open(DATA_FILE)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	decoder := gob.NewDecoder(f)
	var data string
	err = decoder.Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}
