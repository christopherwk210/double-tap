package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/schollz/progressbar/v3"
)

type TargetConfig struct {
	Version  string
	Platform string
	Arch     string
}

func exists(dir string) bool {
	_, err := os.Lstat(dir)
	return err == nil
}

func ensureValidProject(devMode bool) (targetConfig TargetConfig, executablePath string, asarMode bool, packagePath string) {
	platform, _ := detectSystem()

	// Get binary path
	ex, err := os.Executable()
	if err != nil {
		lprintPanic(err, "could not find executable directory")
	}

	if devMode {
		workingDirectory, err := os.Getwd()
		if err != nil {
			lprintPanic(err, "could not get working dev directory")
		}
		ex = filepath.Join(workingDirectory, "temp")
	} else {
		if platform == "darwin" {
			ex = filepath.Join(ex, "../../../../")
		} else {
			ex = filepath.Dir(ex)
		}
	}

	resourcesPath := filepath.Join(ex, "resources")
	appPath := filepath.Join(resourcesPath, "app")
	asarPath := filepath.Join(resourcesPath, "app.asar")
	targetConfigPath := filepath.Join(resourcesPath, "target.json")

	// Check resources path
	if !exists(filepath.Join(ex, "resources")) {
		lprintPanicIdentical("could not find resources directory")
	}

	asarMode = false
	packagePath = appPath
	if exists(asarPath) {
		asarMode = true
		packagePath = asarPath
	} else if !exists(appPath) {
		lprintPanicIdentical("could not locate application files")
	}

	targetConfig = readConfig(targetConfigPath)

	return targetConfig, ex, asarMode, packagePath
}

func readConfig(targetConfigPath string) TargetConfig {
	targetConfigBlob, err := os.ReadFile(targetConfigPath)
	if err != nil {
		lprintPanic(err, "could not read target config")
	}

	var targetConfig TargetConfig
	err = json.Unmarshal(targetConfigBlob, &targetConfig)
	if err != nil {
		lprintPanic(err, "could not parse target config")
	}

	return targetConfig
}

func ensureWorkingDirectory() (workingDir string, tempDir string, buildsDir string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		lprintPanicIdentical("could not find user home directory")
	}

	workingDir = filepath.Join(homeDir, ".double-tap")
	if os.MkdirAll(workingDir, 0777) != nil {
		lprintPanicIdentical("could not create working directory")
	}

	tempDir = filepath.Join(workingDir, "temp")
	if os.MkdirAll(tempDir, 0777) != nil {
		lprintPanicIdentical("could not create temp directory")
	}

	buildsDir = filepath.Join(workingDir, "builds")
	if os.MkdirAll(buildsDir, 0777) != nil {
		lprintPanicIdentical("could not create builds directory")
	}

	return workingDir, tempDir, buildsDir
}

func downloadFile(filepath string, url string, message string) (err error) {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	fmt.Println(message)
	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"downloading",
	)
	defer bar.Close()

	// Writer the body to file
	_, err = io.Copy(io.MultiWriter(out, bar), resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func detectSystem() (os string, arch string) {
	os = runtime.GOOS
	arch = runtime.GOARCH

	if os != "windows" && os != "darwin" && os != "linux" {
		os = "linux"
	}

	if os == "windows" {
		os = "win32"
	}

	if arch == "amd64" {
		arch = "x64"
	}
	if arch == "386" {
		arch = "ia32"
	}

	return os, arch
}

func cleanTemp() {
	if os.RemoveAll(tempPath) != nil {
		lprintPanicIdentical("could not clean temp folder")
	}
}
