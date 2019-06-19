package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 1:1关系
type User struct {
	gorm.Model
	Name     string    `gorm:"type:varchar(64)"`
	Birthday time.Time `gorm:"type:date"`
	Profile  Profile
}

type Profile struct {
	gorm.Model
	Password string `gorm:"type:varchar(1024)"`
	Salt     string `gorm:"type:varchar(64)"`
	Alg      string `gorm:"type:varchar(16)"`
	UserID   uint
}

func main() {
	db, err := gorm.Open("mysql", "root:881019@tcp(localhost:3306)/htgorm?loc=Asia%2FShanghai&charset=utf8mb4&parseTime=True")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer db.Close()
	db.DropTableIfExists("users", "profiles")
	db.AutoMigrate(&User{}, &Profile{})

	profile := Profile{Password: "123!@#", Alg: "text"}

	user := User{Name: "kk", Profile: profile}
	db.Create(&user)

	var u User
	db.First(&u)

	// 关联查询
	var p Profile
	db.Model(&u).Related(&p)
	fmt.Println(u)
	fmt.Println(p)
}
