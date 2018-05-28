package controllers

import (
	"github.com/fly-robin/gopub/models"
	"github.com/fly-robin/gopub/utils"
	"github.com/astaxie/beego"
)

type ProjectController struct {
	BaseController
}


//所有项目列表
func (c *ProjectController) ProjectPanel() {
	c.Prepare()
	projectDetail, err := models.GetProjectById(c.currentProject)
	utils.CheckError(err)
	projectPath := projectDetail.RepositoryPath
	//获取当前所在git相关信息
	gitInfo := utils.GetGitInfo(projectPath)
	c.Data["git_branch"] = gitInfo.Branch
	c.Data["git_tag"] = gitInfo.Tag
	c.Data["project_detail"] = projectDetail
	c.TplName = "project/detail.html"
	c.Render()
}

func getGitInfo(path string) {

}


