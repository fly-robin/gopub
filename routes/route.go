package routes

import (
	"github.com/astaxie/beego"
	"github.com/fly-robin/gopub/controllers"
)

func init()  {
	beego.SetStaticPath("/static", "static")
	//首页
	beego.Router("/", &controllers.IndexController{}, "*:Index")
	//项目列表
	beego.Router("/projectList", &controllers.ProjectController{}, "*:ProjectList")
	//
	beego.Router("/tags", &controllers.GitTagController{}, "*:List")
}