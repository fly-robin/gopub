package utils

import (
	"regexp"
	"strings"
)

//
type GitInfo struct {
	Branch GitBranch //分支信息
	Tag GitTag
}

type GitBranch struct{
	Origin string
	Local string
}

type GitTag struct{
	Tags []string
}


func GetGitInfo(projectPath string) GitInfo {
	localBranch, originBranch := GetCurrentBranch(projectPath)
	gitBranch := GitBranch{
		originBranch,
		localBranch,
	}

	gitTags := GetGitTags(projectPath)
	gitTag := GitTag{
		gitTags,
	}

	var gitInfo GitInfo
	gitInfo.Branch = gitBranch
	gitInfo.Tag = gitTag

	return gitInfo
}

//获取当前的分支名
func GetCurrentBranch(projectPath string)  (string, string) {
	command :=  "cd " + projectPath + " && git status"
	stringResult, err := ExecShell(command)
	CheckError(err)
	//本地分支
	localBranch := findLocalBranchFromGitStatus(stringResult)
	//远程分支
	originBranch := findOriginBranchFromGitStatus(stringResult)

	return localBranch, originBranch
}

//从git status结果中解析出本地分支名
func findLocalBranchFromGitStatus(gitStatus string) string {
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
func GetGitTags(projectPath string) []string {
	commandString := "cd " + projectPath + " & git tag -l"
	gitTagsString, err := ExecShell(commandString)
	CheckError(err)
	gitTagsArray := strings.Split(gitTagsString, "\n")

	return gitTagsArray
}

//更新git本本地仓库
func GitPull(projectPath string) string {
	commandString := "cd " + projectPath + " & git pull"
	result, err := ExecShell(commandString)
	CheckError(err)

	return result
}