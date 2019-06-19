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
}

func main() {
	db, err := gorm.Open("mysql", "root:881019@tcp(localhost:3306)/htgorm?charset=utf8mb4&loc=Asia%2FShanghai&parseTime=True")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	defer db.Close()
	// 删除表
	db.DropTable(&User{})

	// 通过表明删除表
	db.DropTable("users")

	// 当表不存在时跳过删除
	db.DropTableIfExists(&User{})
}
