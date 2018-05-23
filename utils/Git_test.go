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
	result := GetCurrentBranch("./")
	fmt.Println(result)
}