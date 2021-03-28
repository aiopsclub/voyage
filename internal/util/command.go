package util

import (
	"bytes"
	"os/exec"
	"strings"
)

func IsCommandExist(cmd string) bool {
	_, err := exec.LookPath(cmd)
	if err != nil {
		return false
	}
	return true
}

func CommandRun(cmdStr string) (stdout string, stderr string, err error) {
	cmdList := strings.Split(cmdStr, " ")
	cmd := exec.Command(cmdList[0], cmdList[1:]...)
	var stdoutBytes, stderrBytes bytes.Buffer
	cmd.Stdout = &stdoutBytes // 标准输出
	cmd.Stderr = &stderrBytes // 标准错误
	err = cmd.Run()
	stdout, stderr = stdoutBytes.String(), stderrBytes.String()
	return
}
