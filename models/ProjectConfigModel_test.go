package models

import (
	"testing"
	"fmt"
	"log"
)

func TestGetProjectById(t *testing.T) {
	pid := 1
	project, err := GetProjectById(pid)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(err)
		fmt.Println(project)
	}

}
/*
func TestGetProjectList(t *testing.T) {
	list := GetProjectList()
	fmt.Println(list)
}*/