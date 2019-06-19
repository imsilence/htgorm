package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// n:m关系
type User struct {
	gorm.Model
	Name     string     `gorm:"type:varchar(64)"`
	Birthday time.Time  `gorm:"type:date"`
	Langs    []Language `gorm:"many2many:user_language"`
}

type Language struct {
	gorm.Model
	Name string `gorm:"type:varchar(64)"`
}

func main() {
	db, err := gorm.Open("mysql", "root:881019@tcp(localhost:3306)/htgorm?loc=Asia%2FShanghai&charset=utf8mb4&parseTime=True")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer db.Close()
	db.DropTableIfExists("users", "languages", "user_language")
	db.AutoMigrate(&User{}, &Language{})

	lang01 := Language{Name: "汉语"}
	lang02 := Language{Name: "英语"}

	user01 := User{Name: "kk", Langs: []Language{lang01, lang02}}
	db.Create(&user01)

	user02 := User{Name: "silence", Langs: []Language{lang01, lang02}}
	db.Create(&user02)

	var u User
	db.First(&u)

	// 关联查询
	var langs []Language
	db.Model(&u).Related(&langs, "Langs")
	fmt.Println(u)
	fmt.Println(langs)
}
