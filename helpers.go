package main

import (
	"os"
	"os/exec"
	"path"
)

func checkDependencies() error {
	_, err := exec.LookPath("ffmpeg")
	if err != nil {
		return err
	}
	return nil
}

func getAbsPath(p string) (string, error) {
	absPath := p
	if path.IsAbs(absPath) {
	} else {
		dir, err := os.Getwd()
		if err != nil {
			return "", err
		}
		absPath = path.Join(dir, p)
	}
	return absPath, nil
}
