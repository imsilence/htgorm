package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID       int       `gorm:"AUTO_INCREMENT"`
	Name     string    `gorm:"size(64); unique; not null;"`
	Password string    `gorm:"size(1024); not null;"`
	Birthday time.Time `gorm:"index"`
	Desc     string    `gorm:"type:text; not null;"`
}

func (u *User) TableName() string {
	return "user"
}

func main() {
	db, err := gorm.Open("mysql", "root:881019@tcp(127.0.0.1:3306)/htgorm?charset=utf8mb4&loc=Asia%2FShanghai&parseTime=True")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer db.Close()

	db.AutoMigrate(&User{})
}
