package main

import (
	"github.com/astaxie/beego"
	"github.com/fly-robin/gopub/controllers"
	_ "github.com/fly-robin/gopub/routes"
)

func main() {
	beego.Router("/ttt", &controllers.GitTagController{})
	beego.Run()
}
