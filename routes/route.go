package routes

import (
	"github.com/astaxie/beego"
	"github.com/fly-robin/gopub/controllers"
)

func init()  {
	beego.Router("/", &controllers.IndexController{}, "*:Index")

	beego.Router("/projectList", &controllers.ProjectController{}, "*:ProjectList")

	beego.Router("/tags", &controllers.GitTagController{}, "*:List")
}