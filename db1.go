package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)
func main()  {
	conn,err := sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("连接失败：",err)
	}
	defer conn.Close()
	conn.SetMaxOpenConns(2)

	rows,err := conn.Query("select * from users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	_,err = rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	var (
		id int32
		name string
		age int32
	)

	for rows.Next() {
		err = rows.Scan(&id,&name,&age)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id,name,age)
	}
}
