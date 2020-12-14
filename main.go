package main

import (
	_ "dough_go/routers"
	"fmt"

	"github.com/astaxie/beego"
)

func main() {
	str := beego.AppConfig.String("mysqluser")
	fmt.Printf("MyConfig::%v\n", str)
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
