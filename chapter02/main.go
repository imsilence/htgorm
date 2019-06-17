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
}

func main() {
	db, err := gorm.Open("mysql", "root:881019@tcp(127.0.0.1:3306)/htgorm?loc=Asia%2FShanghai&charset=utf8mb4")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	defer db.Close()

	// 自动迁移表
	db.AutoMigrate(&User{})
}