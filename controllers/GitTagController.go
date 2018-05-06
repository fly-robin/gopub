package controllers

import (
	"github.com/fly-robin/gopub/utils"
)

type GitTagController struct {
	BaseController
}

//当前项目标签列表
func (c *GitTagController) List() {
	//c.Data["json"] = map[string]string{"name": "xiaoxu"}
	//c.ServeJSON()
	utils.GetGitTags("./")
}
