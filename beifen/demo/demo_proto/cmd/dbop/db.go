package main

import (
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/biz/dal"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dal.Init()
	// mysql.DB.Create(&model.User{Email: "huqiuyue2021@163.com", Password: "asdfhx123"})
	// mysql.DB.Model(&model.User{}).Where("email=?", "huqiuyue2021@163.com").Update("password", "123456")

	// var row model.User
	// mysql.DB.Model(&model.User{}).Where("email=?", "huqiuyue2021@163.com").First(&row)

	// fmt.Printf("row:%+v\n", row)
	// mysql.DB.Unscoped().Where("email=?", "huqiuyue2021@163.com").Delete(&model.User{})
}
