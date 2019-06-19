package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name     string
	Password string
	Birthday time.Time
	Desc     string
	Status   int
}

type Addr struct {
	gorm.Model
	City   string
	Street string
	UserId uint
}

func main() {
	db, err := gorm.Open("mysql", "root:881019@tcp(localhost:3306)/htgorm?charset=utf8mb4&loc=Asia%2FShanghai&parseTime=True")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	defer db.Close()

	db.CreateTable(&Addr{})
	// 添加外键
	db.Model(&Addr{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
}
