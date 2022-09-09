package main

import (
	"os"
	"path/filepath"
)

func checkBuildExists(buildDir string, targetConfig TargetConfig) bool {
	buildPath, executableName := getBuildPath(buildDir, targetConfig)

	return exists(filepath.Join(buildPath, executableName))
}

func getBuildPath(buildDir string, targetConfig TargetConfig) (buildPath string, executableName string) {
	buildPath = filepath.Join(
		buildDir,
		targetConfig.Platform,
		"_"+targetConfig.Arch,
		"_"+targetConfig.Version,
	)

	switch targetConfig.Platform {
	case "win32":
		executableName = "electron.exe"
	case "linux":
		executableName = "electron"
	case "darwin":
		executableName = "Electron.app"
	}

	return buildPath, executableName
}

func createBuildPath(buildDir string, targetConfig TargetConfig) error {
	buildPath, _ := getBuildPath(buildDir, targetConfig)
	return os.MkdirAll(buildPath, 0777)
}
