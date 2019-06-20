package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name     string
	Password string `gorm:"default:'123!@#QWE'"` //设置默认值
}

func main() {
	db, err := gorm.Open("mysql", "root:881019@tcp(127.0.0.1:3306)/htgorm?charset=utf8mb4&loc=Local&parseTime=True")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer db.Close()
	// db.DropTableIfExists("users")
	// db.AutoMigrate(&User{})
	// for i := 0; i < 10; i++ {
	// 	u := User{Name: fmt.Sprintf("kk_%d", i)}
	// 	if err := db.Create(&u).Error; err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	u = User{Name: fmt.Sprintf("%d_silence", i)}
	// 	if err := db.Create(&u).Error; err != nil {
	// 		fmt.Println(err)
	// 	}
	// }

	db.Model(&User{}).Where("name like ?", "%3%").Update("password", "!@#!@#!@#")

	// 不会执行Callback更新update_at字段
	db.Table("users").Where("name like ?", "%4%").UpdateColumn("password", "123123123")
	db.Model(&User{}).Where("name like ?", "%5%").UpdateColumns(User{Name: "xxxxxxxxx", Password : "123123123"})


	// 只会更新指定字段
	db.Table("users").Where("name like ?", "%6%").Updates(map[string]interface{}{"password" : "xxxx", "name" : "yyyy"})

	// 只会更新更改和非零值的字段
	db.Model(&User{}).Where("name like ?", "%7%").Updates(User{Name: "cccc", Password : "aaaa"})

}
