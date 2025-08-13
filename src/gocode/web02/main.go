package main

//gorm
import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Users struct {
	Id    uint
	Name  string
	Email string
}

func main() {
	//连接mysql数据库 ,登录用户：root，密码cssao888，IP：127.0.0.1:3307，数据库：mygodb
	db, err := gorm.Open("mysql", "root:cssao888@(127.0.0.1:3307)/mygodb?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err) //当程序执行到 panic(v) 时，当前函数会立即停止执行后续代码。
	}
	defer db.Close()

	// 自动迁移GORM 的一个强大功能。AutoMigrate 会检查数据库中是否存在与 Users 结构体对应的表。
	// 如果表不存在，则创建它。
	// 如果表已存在，它会尝试添加缺少的列（但不会删除或修改现有列）。注意：生产环境慎用，因为它可能导致数据丢失或结构不一致。
	db.AutoMigrate(&Users{})

	u1 := Users{3, "七米", "zhaore4415@153.com"}
	u2 := Users{4, "沙河娜扎", "23123@344.com"}
	// 创建记录
	db.Create(&u1)
	db.Create(&u2)
	// 查询
	var u = new(Users)
	db.First(u)
	fmt.Printf("%#v\n", u)

	var uu Users
	db.Find(&uu, "name=?", "沙河娜扎")
	fmt.Printf("%#v\n", uu)

	// 更新
	db.Model(&u).Update("name", "沙河娜扎222")
	// 删除
	//db.Delete(&u)
}
