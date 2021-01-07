package main

import (
	"fmt"
	"log"
	"net/http"
)

func main()  {
	http.HandleFunc("/",test)
	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		log.Fatal("启动失败:",err)
	}
}

func test(w http.ResponseWriter,r *http.Request){
	w.Header().Add("Server","Purelight-Server-01")
	w.WriteHeader(203)
	fmt.Fprintf(w,"hello,world")
}