package main

import (
	_ "NewsWeb/routers"
	"github.com/astaxie/beego"
	_ "NewsWeb/models"
)

func main() {
	beego.Run()
}

