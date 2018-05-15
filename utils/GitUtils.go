package utils

import (
	"os/exec"
	"github.com/astaxie/beego/logs"
)

//
type GitInfo struct {
	branch string //当前所在分支

}

func GetGitTags(projectPath string) string {
	var stringResult string
	commandString := "cd " + projectPath + " & git tag -l"
	cmdResult, err := exec.Command("bash", "-c", commandString).Output()
	if err != nil {
		logs.Error(err)
	}
	stringResult = string(cmdResult)

	return stringResult
}