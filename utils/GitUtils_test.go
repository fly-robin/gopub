package utils

import (
	"testing"
	"fmt"
)

func TestGetGitTags(t *testing.T) {
	result := GetGitTags("./")
	fmt.Println(result)
}