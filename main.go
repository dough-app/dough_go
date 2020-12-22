package main

import (
	_ "dough_go/routers"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type User struct {
	Id   int
	Name string `orm:"size(100)"`
}

func init() {
	// set default database
	_ = orm.RegisterDataBase("default", "mysql", "root:Budengyu1.@tcp(localhost:3306)/dough_app?charset=utf8mb4&parseTime=True", 30)
	// register model
	orm.RegisterModel(new(User))
	// create table
	_ = orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()

	user := User{Name: "slene"}

	// insert
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// update
	user.Name = "astaxie"
	num, err := o.Update(&user)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// read one
	u := User{Id: user.Id}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)

	// delete
	num, err = o.Delete(&u)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	//str := beego.AppConfig.String("mysqluser")
	//fmt.Printf("MyConfig::%v\n", str)
	//if beego.BConfig.RunMode == "dev" {
	//	beego.BConfig.WebConfig.DirectoryIndex = true
	//	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	//}
	//beego.Run()
}
