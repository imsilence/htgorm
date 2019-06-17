package main

import (
	"time"
	"fmt"
	"os"

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
	// 获取数据库连接
	db, err := gorm.Open("mysql", "root:881019@tcp(127.0.0.1:3306)/htgorm?charset=utf8mb4&loc=Asia%2FShanghai")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	// 延迟关闭数据库连接
	defer db.Close()

	// 自动迁移数据库
	db.AutoMigrate(&User{})

	// 插入数据
	for i :=0; i < 10; i++ {
		db.Create(&User{Name: fmt.Sprintf("kk_%d", i), Password : "123!@#"})
	}

	// 按主键获取数据
	var u User
	db.First(&u, 1)
	fmt.Println(u)

	// 按条件获取数据
	var u2 User
	db.First(&u2, "name = ?", "kk_7")
	fmt.Println(u2)

	// 更新数据
	db.Model(u2).Update("password", "!@#QWE!@#")
	fmt.Println(u2)

	// 删除数据
	db.Delete(&u2)
	fmt.Println(u2)
}