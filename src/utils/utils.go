package utils

import (
	"os"
	"os/exec"
	"runtime"
)

func Pwd() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return dir
}

func OpenDir(path string) error {
	switch runtime.GOOS {
	case "windows":
		return exec.Command("explorer", path).Start()
	case "darwin":
		return exec.Command("open", path).Start()
	case "linux":
		return exec.Command("xdg-open", path).Start()
	}
	return nil
}
