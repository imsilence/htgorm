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

	var us []User
	if err := db.Not(&User{Name: "kk_6"}).Find(&us).Error; err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(len(us))
	}

	if err := db.Not("name", "kk_6").Find(&us).Error; err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(len(us))
	}

	if err := db.Not("name = ?", "kk_6").Find(&us).Error; err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(len(us))
	}

	if err := db.Not("name", []string{"kk_4", "kk_6"}).Find(&us).Error; err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(len(us))
	}

	if err := db.Not("name in (?)", []string{"kk_4", "kk_6"}).Find(&us).Error; err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(len(us))
	}

	if err := db.Not([]int{1, 3, 5, 7, 9}).Find(&us).Error; err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(len(us))
	}
}
