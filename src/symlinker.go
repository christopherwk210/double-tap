package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

func createSymLinks(buildDir string, targetConfig TargetConfig, asarMode bool, packagePath string) {
	buildPath, executableName := getBuildPath(buildDir, targetConfig)

	var targetResourcesPath string
	switch targetConfig.Platform {
	case "linux":
		fallthrough
	case "win32":
		targetResourcesPath = filepath.Join(buildPath, "resources")
	case "darwin":
		targetResourcesPath = filepath.Join(buildPath, executableName, "Contents", "Resources")
	}

	var symlinkTarget string
	if asarMode {
		symlinkTarget = filepath.Join(targetResourcesPath, "app.asar")
	} else {
		symlinkTarget = filepath.Join(targetResourcesPath, "app")
	}

	if exists(symlinkTarget) {
		if err := os.Remove(symlinkTarget); err != nil {
			lprintPanic(err, "failed to remove existing symlink")
		}
	}

	if err := os.Symlink(packagePath, symlinkTarget); err != nil {
		lprintPanic(err, "could not create symlink")
	}
}

func runElectron(buildDir string, targetConfig TargetConfig) {
	buildPath, executableName := getBuildPath(buildDir, targetConfig)
	executablePath := filepath.Join(buildPath, executableName)
	var cmd *exec.Cmd
	if targetConfig.Platform == "win32" {
		cmd = exec.Command(executablePath)
	} else {
		cmd = exec.Command("open", executablePath)
	}

	cmd.Env = append(cmd.Env, "ELECTRON_FORCE_IS_PACKAGED=true")
	err := cmd.Start()
	if err != nil {
		lprintPanic(err, "failed to run electron")
	}
}
