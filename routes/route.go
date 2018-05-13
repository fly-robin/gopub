package routes

import (
	"github.com/astaxie/beego"
	"github.com/fly-robin/gopub/controllers"
)

func init() {
	beego.SetStaticPath("/static", "static")
	//首页
	beego.Router("/", &controllers.IndexController{}, "*:Index")
	//项目列表
	//beego.Router("/project/list", &controllers.ProjectController{}, "*:ProjectList")
	beego.Router("/project/detail", &controllers.ProjectController{}, "*:ProjectPanel")
	//
	beego.Router("/tags", &controllers.GitTagController{}, "*:List")
}