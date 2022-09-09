package main

import (
	"path/filepath"
)

var baseDownloadURL string
var targetConfig TargetConfig
var artifactName string

var executablePath string
var zipPath string
var asarMode bool
var packagePath string

var tempPath string
var buildsPath string

func init() {
	os, arch := detectSystem()
	targetConfig, executablePath, asarMode, packagePath = ensureValidProject()
	_, tempPath, buildsPath = ensureWorkingDirectory()

	if targetConfig.Version == "" {
		lprintPanicIdentical("must specify electron version")
	}

	if targetConfig.Arch == "" {
		targetConfig.Arch = arch
	}

	if targetConfig.Platform == "" {
		targetConfig.Platform = os
	}

	artifactName = "electron-v" +
		targetConfig.Version + "-" +
		targetConfig.Platform + "-" +
		targetConfig.Arch + ".zip"

	zipPath = filepath.Join(tempPath, artifactName)
	baseDownloadURL = "https://github.com/electron/electron/releases/download/v"

	setupLogger(filepath.Join(executablePath, "log.txt"))
}

func main() {
	getArtifact()

	createSymLinks(buildsPath, targetConfig, asarMode, packagePath)
	runElectron(buildsPath, targetConfig)

	closeLogFile()
	cleanTemp()
}

func getArtifact() {
	if !checkBuildExists(buildsPath, targetConfig) {
		if err := createBuildPath(buildsPath, targetConfig); err != nil {
			lprintPanic(err, "could not create output build path")
		}

		downloadURL := baseDownloadURL +
			targetConfig.Version +
			"/" +
			artifactName

		err := downloadFile(zipPath, downloadURL, "Downloading artifact...")
		if err != nil {
			lprintPanic(err, "could not download artifact")
		}

		println("Unzipping artifact...")
		buildPath, _ := getBuildPath(buildsPath, targetConfig)

		if targetConfig.Platform == "darwin" {
			if err := nativeUnzip(zipPath, buildPath); err != nil {
				lprintPanic(err, "could not unzip artifact")
			}
		} else {
			if err := unzipSource(zipPath, buildPath); err != nil {
				lprintPanic(err, "could not unzip artifact")
			}
		}

		println("Done.")
	}
}
