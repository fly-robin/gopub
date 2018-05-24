package utils

import (
	"regexp"
	"fmt"
)

//
type GitInfo struct {
	branch string //当前所在分支

}

//获取当前的分支名
func GetCurrentBranch(projectPath string)  (string, string) {
	cmommand :=  "cd " + projectPath + " & git status"
	stringResult, err := ExecShell(cmommand)
	CheckError(err)
	localBranch := findLocalBranchFromGitStatus(stringResult)
	originBranch := findOriginBranchFromGitStatus(stringResult)

	return localBranch, originBranch
}

//从git status结果中解析出本地分支名
func findLocalBranchFromGitStatus(gitStatus string) string {
	fmt.Println(gitStatus)
	reg, err := regexp.Compile("branch (.*)\n");
	subMatch := reg.FindStringSubmatch(gitStatus)
	CheckError(err)
	var branch string
	if len(subMatch) > 0 {
		branch = subMatch[1]
	}

	return branch
}

//解析出对应的远程分支
func findOriginBranchFromGitStatus(gitStatus string) string {
	fmt.Println(gitStatus)
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