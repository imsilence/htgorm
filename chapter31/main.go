package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 1:n关系
type User struct {
	gorm.Model
	Name     string    `gorm:"type:varchar(64)"`
	Birthday time.Time `gorm:"type:date"`
	Addrs    []Address
}

type Address struct {
	gorm.Model
	City   string
	Street string
	UserId uint
}

func main() {
	db, err := gorm.Open("mysql", "root:881019@tcp(localhost:3306)/htgorm?loc=Asia%2FShanghai&charset=utf8mb4&parseTime=True")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer db.Close()
	db.DropTableIfExists("users", "addresses")
	db.AutoMigrate(&User{}, &Address{})

	addr01 := Address{City: "北京", Street: "海淀"}
	addr02 := Address{City: "西安", Street: "高新"}

	user := User{Name: "kk", Addrs: []Address{addr01, addr02}}
	db.Create(&user)

	db = db.Table("users").Select("users.name, addresses.city, addresses.street")
	rows, _ := db.Joins("left join addresses on users.id = addresses.user_id").Rows()
	for rows.Next() {
		var name, city, street string
		rows.Scan(&name, &city, &street)
		fmt.Println(name, city, street)
	}

}
