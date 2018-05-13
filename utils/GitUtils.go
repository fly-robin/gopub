package utils

import (
	"os/exec"
	"github.com/astaxie/beego/logs"
)

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