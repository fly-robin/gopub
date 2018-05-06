package models

import "github.com/astaxie/beego/orm"

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

func GetProjectList() []*ProjectConfig {
	projects := make([]*ProjectConfig, 0)
	orm.NewOrm().QueryTable(TableName("project")).All(&projects)
	//num, err := o.quer
	return projects
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
