package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sync"
	"time"
)

var wg2 sync.WaitGroup

func main()  {

	conn,err := sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("连接失败：",err)
	}
	conn.SetMaxOpenConns(1)

	defer conn.Close()

	wg2.Add(10)
	for i := 0;i < 10;i++ {
		go query(conn,i)
	}
	wg2.Wait()

	insert(conn)

	fmt.Println("Done")
}

func query(conn *sql.DB,i int)  {
	var (
		id int
		name string
		age string
	)
	rows,err := conn.Query("select id,name,age from users where id = ?",1)
	if err != nil {
		fmt.Println("查询失败：",err)
	}
	defer rows.Close()

	time.Sleep(time.Second)

	for rows.Next()  {
		err = rows.Scan(&id,&name,&age)
		if err != nil {
			fmt.Println("取值失败：",err)
		}
		fmt.Println(i,time.Now().Format("2006-01-02 15:04:05"),id,name,age)
		go q(conn,i)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println("错误：",err)
	}

	wg2.Done()
}

func q(conn *sql.DB,i int)  {
	var name string
	var has bool
	err := conn.QueryRow("select name from users where id = ?",80).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			has = false
		}else {
			log.Fatal(err)
		}
	}else {
		has = true
	}
	fmt.Println(has,name)
}

func insert(conn *sql.DB)  {

	tx,err := conn.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt,err := tx.Prepare("insert into users(name,age) values(?,?)")
	if err != nil {
		log.Fatal(err)
	}

	res,err := stmt.Exec("golang",10)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	lastId,err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	affectedRows,err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ID:",lastId,"AffectedRows:",affectedRows)

	_,err = tx.Exec("delete from users where id = ?",lastId)
	if err != nil {
		tx.Rollback()
		fmt.Println("delete err : ",err)
	}else {
		fmt.Println("删除成功:",lastId)
	}

	tx.Commit()
}