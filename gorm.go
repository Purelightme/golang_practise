package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"purelight/orm/model"
)

func main()  {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:root@tcp(localhost:3306)/test?parseTime=true&loc=Local",
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&model.Teacher{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&model.Student{})
	if err != nil {
		log.Fatal(err)
	}

	user := model.Student{}
	tx := db.Create(&model.Student{Name:"张三",Teacher:model.Teacher{Name:"张老师",Age:30}})
	if tx.Error != nil {
		log.Fatal(tx.Error)
	}

	tx.Scan(&user)

	fmt.Println(user)
}
