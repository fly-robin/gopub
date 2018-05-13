package controllers

//发布系统主页
type IndexController struct {
	BaseController
}


//主页面承载体
func (this *IndexController) Index() {
	//fmt.Println(beego.BConfig.AppName)
	this.Data["description"] = "这是一个基于beego开发的代码发布系统，开发此工主要是为了方便php代码上线部署。目前v0.1版本支持基于git版本控制下的对tag、branch进行增量代码发布，以及增量对应的代码回滚。除此之外还做了代码上线历史记录。"
	this.TplName = "index/index.html"
	this.Data["current_project"] = 0
	this.Render()
}