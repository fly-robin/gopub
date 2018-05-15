package controllers

import (
	"github.com/astaxie/beego"
	"github.com/fly-robin/gopub/models"
)

type BaseController struct {
	beego.Controller
	currentProject int
	currentProjectName string
}

func (c *BaseController) init()  {
}

func (this *BaseController) Prepare() {
	//站点公共标题
	this.Data["SITE_NAME"] = beego.AppConfig.String("site.name")
	//项目列表
	projects, num := models.GetProjectList()
	this.Data["projects_list"] = projects
	this.Data["projects_num"] = num
	var err error
	if this.currentProject, err = this.GetInt("current_project"); err != nil {
		this.currentProject = projects[0].Id
		this.currentProjectName = projects[0].Name
	}
	this.Data["current_project"] = this.currentProject
}