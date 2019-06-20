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

// 定义回调函数，在创建之前执行
func (u *User) BeforeCreate(scope *gorm.Scope) error {
	fmt.Println("before create")
	return nil
}

func main() {
	db, err := gorm.Open("mysql", "root:881019@tcp(127.0.0.1:3306)/htgorm?charset=utf8mb4&loc=Local&parseTime=True")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer db.Close()
	db.AutoMigrate(&User{})

	u := User{Name: "kk"}
	fmt.Println(db.NewRecord(u)) // 判断对象是否已经被创建

	db.Create(&u)
	fmt.Println(db.NewRecord(u))

}
