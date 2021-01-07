package main

import (
	"fmt"
	"reflect"
)

func main(){
	name := "小米"
	value := reflect.ValueOf(&name)
	value.Elem().Set(reflect.ValueOf("华为"))
	fmt.Println(name)
}
