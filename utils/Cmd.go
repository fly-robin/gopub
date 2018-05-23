package utils

import "os/exec"

func ExecShell(command string) (string, error) {
	result, err := exec.Command("bash", "-c", command).Output()
	stringResult := string(result)

	return stringResult, err
}