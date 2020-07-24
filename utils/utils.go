package utils

import (
	"fmt"
	"os/exec"
	"runtime"
)

//CheckPortUsedForTCP 检查TCP端口是否已被占用
func CheckPortUsedForTCP(port string) (used bool, err error) {

	var (
		command     string
		commandArgs []string
	)

	switch runtime.GOOS {
	case "linux":

		command = "sh"
		commandArgs = append(commandArgs, "-c", fmt.Sprintf("netstat -anpt | grep %s", port))
	case "windows":

		command = "cmd"
		commandArgs = append(commandArgs, "/c", fmt.Sprintf("netstat -ano -p tcp | findstr %s", port))
	case "darwin":

		command = "sh"
		commandArgs = append(commandArgs, "-c", fmt.Sprintf("lsof -i:%s", port))
	}

	execOutput, err := exec.Command(command, commandArgs...).CombinedOutput()

	if len(execOutput) > 0 {

		used = true
	}

	return
}
