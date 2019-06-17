package main

import (
	"os"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name string
	Password string
	Birthday time.Time
	Desc string
	Status int
}

func main() {
	db, err := gorm.Open("mysql", "root:881019@tcp(localhost:3306)/htgorm?charset=utf8mb4&loc=Asia%2FShanghai")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	defer db.Close()

	// 修改列的类型
	db.Model(&User{}).ModifyColumn("birthday", "datetime")
}