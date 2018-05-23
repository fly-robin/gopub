package utils

import (
	"regexp"
)

//
type GitInfo struct {
	branch string //当前所在分支

}

//获取当前的分支名
func GetCurrentBranch(projectPath string)  string {
	cmommand :=  "cd " + projectPath + " & git status"
	result, err := ExecShell(cmommand)
	CheckError(err)
	stringResult := string(result)
	currentBranch := findBranchFromGitStatus(stringResult)

	return currentBranch
}

//从git status结果中解析出分支名
func findBranchFromGitStatus(gitStatus string) string {
	reg, err := regexp.Compile("\\'(.*)\\'");
	subMatch := reg.FindStringSubmatch(gitStatus)
	CheckError(err)
	var branch string
	if len(subMatch) > 0 {
		branch = subMatch[1]
	}

	return branch
}


//获取当前工程的所有tag
func GetGitTags(projectPath string) string {
	commandString := "cd " + projectPath + " & git tag -l"
	result, err := ExecShell(commandString)
	CheckError(err)

	return result
}

//更新git本本地仓库
func GitPull(projectPath string) string {
	commandString := "cd " + projectPath + " & git pull"
	result, err := ExecShell(commandString)
	CheckError(err)

	return result
}