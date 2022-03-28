package util

import (
	"os"
	"syscall"
)

func GetExecDirectory() string {
	file, err := os.Getwd()
	if err == nil {
		return file + "/"
	}
	return ""
}

func CheckProcessExist(pid int) bool {
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}

	if err = process.Signal(syscall.Signal(0)); err != nil {
		return false
	}

	return true
}
