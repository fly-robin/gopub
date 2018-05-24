package utils

import (
	"testing"
	"fmt"
)

func TestGetGitTags(t *testing.T) {
	result := GetGitTags("./")
	fmt.Println(result)
}

func TestGetCurrentBranch(t *testing.T) {
	local, origin := GetCurrentBranch("/d/jr/sina_finance")
	fmt.Println(local, origin)
}