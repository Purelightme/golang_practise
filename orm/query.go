package orm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"purelight/orm/model"
)

var DB *gorm.DB
var err error

func Init()  {
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN: "root:root@tcp(localhost:3306)/test?parseTime=true&loc=Local",
	}), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}

func SimpleQuery()  {
	Init()

	//user := model.User{}
	//tx := DB.Where("name = ?","golang").Take(&user)
	//if tx.Error != nil {
	//	if errors.Is(tx.Error,gorm.ErrRecordNotFound) {
	//		fmt.Println("没找到记录")
	//	}
	//}
	//fmt.Println(user.ID)

	//type Result struct {
	//	Total int64
	//}
	//res := Result{}
	//DB.Table("users").Select("count(*) as Total").Find(&res)
	//if DB.Error != nil {
	//	log.Fatal(DB.Error)
	//}
	//fmt.Println(res.Total)

	//users := make([]model.User,1)
	//DB.Table("users").
	//	Joins("left join user_profiles on user_profiles.user_id = users.id",).
	//	Select("users.Name","users.ID","users.Age","Coin").
	//	Find(&users)
	//printUsers(users)

	//profiles := make([]model.User,1)
	//DB.Preload("UserProfile").Find(&profiles)
	//printUsers(profiles)

	profiles := make([]model.UserProfile,1)
	DB.Preload("User").Find(&profiles)
	printProfiles(profiles)
}

//func printUsers(users []model.User)  {
//	fmt.Println("查询结果：")
//	for _,v := range users {
//		fmt.Println(v.ID,v.Name,v.Age,v.UserProfile.Coin)
//	}
//}

func printProfiles(profiles []model.UserProfile)  {
	fmt.Println("查询结果：")
	for _,v := range profiles {
		fmt.Println(v.ID,v.Coin,v.User.Name)
	}
}
