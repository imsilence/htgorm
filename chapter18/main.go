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
	}

	// 获取第一条数据
	var u01 User
	db.First(&u01)
	fmt.Println(u01)

	// 获取最后一条数据
	var u02 User
	db.Last(&u02)
	fmt.Println(u02)

	// 按主键获取
	var u03 User
	db.First(&u03, 8)
	fmt.Println(u03)

	// 按条件查询一条
	var u04 User
	db.First(&u04, "name=?", "kk_5")
	fmt.Println(u04)

	// 获取所有数据
	var us []User
	db.Find(&us)
	fmt.Println(us)

	db.Find(&us, "name != ?", "kk_5")
	fmt.Println(us)
}
