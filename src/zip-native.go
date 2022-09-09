package main

import "os/exec"

func nativeUnzip(source string, destination string) error {
	os, _ := detectSystem()

	switch os {
	case "darwin":
		cmd := exec.Command("unzip", source, "-d", destination)
		err := cmd.Run()
		return err
	}

	return nil
}
