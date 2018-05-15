package controllers

import (
	"github.com/fly-robin/gopub/models"
	"github.com/fly-robin/gopub/utils"
)

type ProjectController struct {
	BaseController
}


//所有项目列表
func (c *ProjectController) ProjectPanel() {
	c.Prepare()
	projectDetail, err := models.GetProjectById(c.currentProject)
	utils.CheckError(err)

	c.Data["project_detail"] = projectDetail
	c.TplName = "project/detail.html"
	c.Render()
}

func getGitInfo(path string) {

}


