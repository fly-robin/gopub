package controllers

import (
)

type ProjectController struct {
	BaseController
}


//所有项目列表
func (c *ProjectController) ProjectPanel() {
	c.Prepare()
	c.TplName = "project/detail.html"
	c.Render()
}


