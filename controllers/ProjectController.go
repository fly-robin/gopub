package controllers

import (
	"github.com/fly-robin/gopub/models"
	"fmt"
)

type ProjectController struct {
	BaseController
}

func (this ProjectController) ProjectList() {
	p, err := models.GetProjectById(1)
	fmt.Println(p)
	fmt.Println(err)
}
