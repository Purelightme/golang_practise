package main

import (
	"context"
	"fmt"
	"time"
)

type User struct {
	name string
	age int
}

func main()  {

	user := User{name:"scl",age:20}

	var cancel func()
	ctx := context.WithValue(context.Background(),"user",user)
	ctx,cancel = context.WithCancel(ctx)

	go do(ctx)

	time.Sleep(time.Second * 3)

	cancel()
}

func do(ctx context.Context)  {
	u := ctx.Value("user").(User)
	fmt.Println("name is ",u.name," age is ",u.age)
}