package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/fly-robin/gopub/utils"
)

type ProjectConfig struct {
	Id             int
	Name           string
	RepositoryPath string
	CTime          string
	UTime          string
}

func (this *ProjectConfig) TableName() string {
	return TableName("project")
}

func GetProjectList() ([]ProjectConfig, int64) {
	projects := make([]ProjectConfig, 0)
	query := orm.NewOrm().QueryTable(TableName("project"))
	//查询列表
	query.All(&projects)
	//项目数
	count, err := query.Count()
	utils.CheckError(err)

	return projects, count
}

//获取项目信息
func GetProjectById(id int) (*ProjectConfig, error) {
	project := &ProjectConfig{
		Id: id,
	}
	err := orm.NewOrm().Read(project)
	if err != nil {
		return nil, err
	}

	return project, err
}
