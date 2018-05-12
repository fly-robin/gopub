package controllers

import (
	"github.com/fly-robin/gopub/models"
)

type ProjectController struct {
	BaseController
}


//所有项目列表
func (c *ProjectController) ProjectList() {
	projects, num := models.GetProjectList()
	c.Data["projects_list"] = projects
	c.Data["projects_num"] = num
	c.TplName = "project/list"
	//fmt.Println(projects)
}
