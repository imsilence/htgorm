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

	var count int
	if err := db.Model(&User{}).Count(&count).Error; err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(count)
	}

	if err := db.Model(&User{}).Where("name like ?", "%6%").Count(&count).Error; err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(count)
	}

	if err := db.Table("users").Where("name like ?", "%6%").Count(&count).Error; err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(count)
	}


}
