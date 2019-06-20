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
	db.DropTableIfExists("users")
	db.AutoMigrate(&User{})
	for i := 0; i < 10; i++ {
		u := User{Name: fmt.Sprintf("kk_%d", i)}
		if err := db.Create(&u).Error; err != nil {
			fmt.Println(err)
		}
		u = User{Name: fmt.Sprintf("%d_silence", i)}
		if err := db.Create(&u).Error; err != nil {
			fmt.Println(err)
		}
	}

	var us01, us02 []User
	// 分别查询按name降序、升序排列结果
	if err := db.Order("name desc").Find(&us01).Order("name asc", true).Find(&us02).Error; err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(us01)
		fmt.Println(us02)
	}
}
