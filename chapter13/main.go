package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User 属于 Profile， 外键为ProfileID
type User struct {
	gorm.Model
	Name      string    `gorm:"type:varchar(64)"`
	Birthday  time.Time `gorm:"type:date"`
	Profile   Profile
	ProfileID uint
}

type Profile struct {
	gorm.Model
	Alg      string `gorm:"type:varchar(16)"`
	Salt     string `gorm:"type:varchar(64)"`
	Password string `gorm:"type:varchar(1024)"`
}

func main() {
	db, err := gorm.Open("mysql", "root:881019@tcp(127.0.0.1:3306)/htgorm?charset=utf8mb4&loc=Asia%2FShanghai&parseTime=True")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer db.Close()
	db.DropTableIfExists("users", "profiles")
	db.AutoMigrate(&User{}, &Profile{})

	profile := Profile{Alg: "text", Salt: "", Password: "123!@#"}
	user := User{Name: "kk", Profile: profile}
	db.Create(&user)

	var u User

	db.First(&u, 1)

	var p Profile

	// 关联查询Profile
	db.Model(&u).Related(&p)

	fmt.Println(u)
	fmt.Println(p)

}
